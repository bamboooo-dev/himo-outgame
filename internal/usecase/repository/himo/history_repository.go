package himo

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// HistoryRepository はインターフェース
type HistoryRepository interface {
	FetchIDsByUser(ctx context.Context, db *gorp.DbMap, user model.User) ([]uint32, error)
	FetchHistoriesByIDs(ctx context.Context, db *gorp.DbMap, user model.User, historyIDs []uint32) ([]model.History, error)
}
