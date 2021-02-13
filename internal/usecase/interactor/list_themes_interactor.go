package interactor

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/bamboooo-dev/himo-outgame/internal/domain/service"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// ListThemeInteractor は自分が作ったお題を見るユースケースを司る構造体
type ListThemeInteractor struct {
	userRepo     himo_repo.UserRepository
	themeRepo    himo_repo.ThemeRepository
	themeService *service.ThemeService
}

// NewListThemeInteractor は ListThemeInteractor のコンストラクタ
func NewListThemeInteractor(r registry.Registry) *ListThemeInteractor {
	return &ListThemeInteractor{
		userRepo:     r.NewUserRepository(),
		themeRepo:    r.NewThemeRepository(),
		themeService: service.NewThemeService(r),
	}
}

// Call は自分が作ったお題を返す関数
func (c *ListThemeInteractor) Call(ctx context.Context, db *gorp.DbMap, userID string) ([]model.Theme, error) {
	user, err := c.userRepo.Find(ctx, db, userID)
	if err != nil {
		return []model.Theme{}, err
	}
	themes, err := c.themeService.List(ctx, db, user)
	if err != nil {
		return []model.Theme{}, err
	}
	return themes, nil
}
