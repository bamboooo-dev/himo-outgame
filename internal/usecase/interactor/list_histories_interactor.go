package interactor

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/bamboooo-dev/himo-outgame/internal/domain/service"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// ListHistoriesInteractor は自分が作ったお題を見るユースケースを司る構造体
type ListHistoriesInteractor struct {
	userRepo       himo_repo.UserRepository
	historyRepo    himo_repo.HistoryRepository
	historyService *service.HistoryService
}

// NewListHistoriesInteractor は ListHistoriesInteractor のコンストラクタ
func NewListHistoriesInteractor(r registry.Registry) *ListHistoriesInteractor {
	return &ListHistoriesInteractor{
		userRepo:       r.NewUserRepository(),
		historyRepo:    r.NewHistoryRepository(),
		historyService: service.NewHistoryService(r),
	}
}

// Call は自分が作ったお題を返す関数
func (c *ListHistoriesInteractor) Call(ctx context.Context, db *gorp.DbMap, userID string) ([]model.History, error) {
	user, err := c.userRepo.Find(ctx, db, userID)
	if err != nil {
		return []model.History{}, err
	}
	histories, err := c.historyService.List(ctx, db, user)
	if err != nil {
		return []model.History{}, err
	}
	return histories, nil
}
