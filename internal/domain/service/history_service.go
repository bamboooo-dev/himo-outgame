package service

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// HistoryService は履歴に関するドメインサービスの構造体
type HistoryService struct {
	userRepo    repo.UserRepository
	historyRepo repo.HistoryRepository
}

// NewHistoryService は HistoryService のコンストラクタ
func NewHistoryService(r registry.Registry) *HistoryService {
	return &HistoryService{
		userRepo:    r.NewUserRepository(),
		historyRepo: r.NewHistoryRepository(),
	}
}

// List は自分が遊んだ履歴を取得する
func (s *HistoryService) List(ctx context.Context, db *gorp.DbMap, user model.User) ([]model.History, error) {
	historyIDs, err := s.historyRepo.FetchIDsByUser(ctx, db, user)
	if err != nil {
		return nil, err
	}

	histories, err := s.historyRepo.FetchHistoriesByIDs(ctx, db, user, historyIDs)
	if err != nil {
		return nil, err
	}

	return histories, nil
}
