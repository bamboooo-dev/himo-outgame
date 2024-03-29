package interactor

import (
	"context"
	"strconv"
	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// UpdateUserInteractor ユーザーを更新するアプリケーションサービス
type UpdateUserInteractor struct {
	userRepo himo_repo.UserRepository
}

// NewUpdateUserInteractor constructs UpdateUserInteractor
func NewUpdateUserInteractor(r registry.Registry) *UpdateUserInteractor {
	return &UpdateUserInteractor{
		userRepo: r.NewUserRepository(),
	}
}

func (u *UpdateUserInteractor) Call(ctx context.Context, db *gorp.DbMap, nickName string, userID string) error {
	intUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return err
	}
	user := model.User{
		Nickname: nickName,
		ID:       intUserID}
	_ , err = u.userRepo.Update(ctx, db, user)
	if err != nil {
		return err
	}
	return nil
}
