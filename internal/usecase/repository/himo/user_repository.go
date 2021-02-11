package himo

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// UserRepository はインターフェース
type UserRepository interface {
	Create(ctx context.Context, db *gorp.DbMap, user model.User) (model.User, error)
}
