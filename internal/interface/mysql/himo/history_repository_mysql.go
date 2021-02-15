package himo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

type historyUserInfo struct {
	HistoryID uint32    `db:"history_id"`
	UserID    uint32    `db:"user_id"`
	Nickname  string    `db:"nickname"`
	CreatedAt time.Time `db:"created_at"`
}

// HistoryRepositoryMysql は HistoryRepository の MySQL 実装
type HistoryRepositoryMysql struct {
	logger *zap.SugaredLogger
}

// NewHistoryRepositoryMysql は HistoryRepositoryMysql のコンストラクタ
func NewHistoryRepositoryMysql(l *zap.SugaredLogger) repo.HistoryRepository {
	return HistoryRepositoryMysql{logger: l}
}

// FetchIDsByUser fetch history ID by a user
func (h HistoryRepositoryMysql) FetchIDsByUser(ctx context.Context, db *gorp.DbMap, user model.User) ([]uint32, error) {

	// TODO: テストする必要あり
	var ids []uint32

	_, err := db.Select(&ids, "SELECT history_id FROM user_histories WHERE user_id = ?", user.ID)
	if err != nil {
		return []uint32{}, err
	}
	return ids, nil
}

// FetchHistoriesByIDs fetches histories by IDs
func (h HistoryRepositoryMysql) FetchHistoriesByIDs(ctx context.Context, db *gorp.DbMap, user model.User, historyIDs []uint32) ([]model.History, error) {
	if len(historyIDs) == 0 {
		return nil, nil
	}
	var daoHistories []historyUserInfo
	args := make([]interface{}, len(historyIDs))
	quarks := make([]string, len(historyIDs))
	for i, historyID := range historyIDs {
		args[i] = historyID
		quarks[i] = "?"
	}
	args = append(args, user.ID)

	_, err := db.Select(&daoHistories,
		fmt.Sprintf("SELECT hu.history_id, hu.user_id, u.nickname, h.created_at "+
			"FROM user_histories AS hu INNER JOIN users AS u ON hu.user_id = u.id INNER JOIN histories AS h ON hu.history_id = h.id "+
			"WHERE hu.history_id IN (%s) AND hu.user_id != ?", strings.Join(quarks, ", ")), args...)
	if err != nil {
		return []model.History{}, err
	}
	histories := make([]model.History, len(historyIDs))
	for i, id := range historyIDs {
		withUsers := []model.User{}
		var createdAt time.Time
		for _, daoHistory := range daoHistories {
			if daoHistory.HistoryID == id {
				withUsers = append(withUsers, model.User{
					ID:       int64(daoHistory.UserID),
					Nickname: daoHistory.Nickname,
				})
				createdAt = daoHistory.CreatedAt
			}
		}
		histories[i] = model.History{
			ID:        id,
			CreatedAt: createdAt,
			WithUsers: withUsers,
		}
	}
	return histories, nil
}
