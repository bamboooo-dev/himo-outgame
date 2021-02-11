package himo

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	dao "github.com/bamboooo-dev/himo-outgame/internal/interface/dao/himo"
	repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// UserRepositoryMysql は UserRepository の MySQL 実装
type UserRepositoryMysql struct {
	logger *zap.SugaredLogger
}

// NewUserRepositoryMysql は UserRepositoryMysql のコンストラクタ
func NewUserRepositoryMysql(l *zap.SugaredLogger) repo.UserRepository {
	return UserRepositoryMysql{logger: l}
}

// Create inserts new user
func (u UserRepositoryMysql) Create(ctx context.Context, db *gorp.DbMap, user model.User) (model.User, error) {
	userDAO := &dao.User{Nickname: user.Nickname}

	err := db.Insert(userDAO)
	if err != nil {
		return model.User{}, err
	}

	user = model.User{Nickname: userDAO.Nickname}
	return user, nil
}