package himo

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// ThemeRepository はインターフェース
type ThemeRepository interface {
	Create(ctx context.Context, db *gorp.DbMap, user model.User, sentence string) (model.Theme, error)
	FetchByUser(ctx context.Context, db *gorp.DbMap, user model.User) ([]model.Theme, error)
}
