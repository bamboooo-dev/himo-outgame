package interactor

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/bamboooo-dev/himo-outgame/internal/domain/service"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// CreateThemeInteractor はお題を作るユースケースを司る構造体
type CreateThemeInteractor struct {
	userRepo     himo_repo.UserRepository
	themeRepo    himo_repo.ThemeRepository
	themeService *service.ThemeService
}

// NewCreateThemeInteractor は CreateThemeInteractor のコンストラクタ
func NewCreateThemeInteractor(r registry.Registry) *CreateThemeInteractor {
	return &CreateThemeInteractor{
		userRepo:     r.NewUserRepository(),
		themeRepo:    r.NewThemeRepository(),
		themeService: service.NewThemeService(r),
	}
}

// Call はお題を作る関数
func (c *CreateThemeInteractor) Call(ctx context.Context, db *gorp.DbMap, userID string, sentence string) (model.Theme, error) {
	user, err := c.userRepo.Find(ctx, db, userID)
	if err != nil {
		return model.Theme{}, err
	}
	theme, err := c.themeService.Create(ctx, db, user, sentence)
	if err != nil {
		return model.Theme{}, err
	}
	return theme, nil
}
