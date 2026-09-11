package http

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	"github.com/bamboooo-dev/himo-outgame/internal/usecase/interactor"
	"github.com/bamboooo-dev/himo-outgame/pkg/grpcmiddleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// UserHandler は Web クライアント向けの JSON API。
// ブラウザからは gRPC を直接叩けないので、UserManager のうちクライアントが使う
// SignUp / UpdateUserName だけを HTTP で提供する。ロジックは gRPC と同じ interactor を使う。
type UserHandler struct {
	logger   *zap.SugaredLogger
	register *interactor.RegisterUserInteractor
	updater  *interactor.UpdateUserInteractor
	db       *gorp.DbMap
}

// NewUserHandler returns UserHandler
func NewUserHandler(l *zap.SugaredLogger, r registry.Registry, db *gorp.DbMap) *UserHandler {
	return &UserHandler{
		logger:   l,
		register: interactor.NewRegistUserInteractor(r),
		updater:  interactor.NewUpdateUserInteractor(r),
		db:       db,
	}
}

// Register は mux にエンドポイントを登録する
func (h *UserHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/signup", cors(h.SignUp))
	mux.HandleFunc("/nickname", cors(h.UpdateNickname))
}

type nicknameRequest struct {
	Nickname string `json:"nickname"`
}

// SignUp は POST /signup {"nickname": "..."} でユーザーを作り JWT を返す
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	req, ok := readNickname(w, r)
	if !ok {
		return
	}

	token, err := h.register.Call(r.Context(), h.db, req.Nickname)
	if err != nil {
		h.logger.Error(err)
		writeError(w, http.StatusInternalServerError, "failed to sign up")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"access_token": token,
		"nickname":     req.Nickname,
	})
}

// UpdateNickname は POST /nickname {"nickname": "..."} (要 Bearer) でニックネームを更新する
func (h *UserHandler) UpdateNickname(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	userID, ok := authenticate(r)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	req, ok := readNickname(w, r)
	if !ok {
		return
	}

	if err := h.updater.Call(r.Context(), h.db, req.Nickname, userID); err != nil {
		h.logger.Error(err)
		writeError(w, http.StatusInternalServerError, "failed to update nickname")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"nickname": req.Nickname})
}

// authenticate は Authorization: Bearer <jwt> を検証して user_id を返す(gRPC の Authenticate と同じ claim)
func authenticate(r *http.Request) (string, bool) {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(strings.ToLower(auth), "bearer ") {
		return "", false
	}
	tokenString := strings.TrimSpace(auth[len("bearer "):])

	claim := grpcmiddleware.AuthClaim{}
	_, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || claim.UserID == "" {
		return "", false
	}
	return claim.UserID, true
}

func readNickname(w http.ResponseWriter, r *http.Request) (nicknameRequest, bool) {
	var req nicknameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return req, false
	}
	req.Nickname = strings.TrimSpace(req.Nickname)
	if req.Nickname == "" {
		writeError(w, http.StatusBadRequest, "nickname is required")
		return req, false
	}
	return req, true
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]interface{}{"code": status, "message": message})
}

// cors は開発時に Web クライアントを別オリジン(例: file:// や別ポート)から動かせるようにする
func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}
