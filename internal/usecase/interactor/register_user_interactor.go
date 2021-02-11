package interactor

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// RegisterUserInteractor ユーザーを登録するアプリケーションサービス
type RegisterUserInteractor struct {
	userRepo himo_repo.UserRepository
}

// NewRegistUserInteractor constructs RegisterUserInteractor
func NewRegistUserInteractor(r registry.Registry) *RegisterUserInteractor {
	return &RegisterUserInteractor{
		userRepo: r.NewUserRepository(),
	}
}

// Call は受け取ったニックネームでユーザーを登録する関数
func (r *RegisterUserInteractor) Call(ctx context.Context, db *gorp.DbMap, nickName string) (string, error) {
	user := model.User{
		Nickname: nickName,
	}
	user, err := r.userRepo.Create(ctx, db, user)
	if err != nil {
		return "", err
	}
	return user.AccessToken, nil
}
