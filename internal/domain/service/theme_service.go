package service

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// ThemeService はお題に関するドメインサービスの構造体
type ThemeService struct {
	userRepo  repo.UserRepository
	themeRepo repo.ThemeRepository
}

// NewThemeService は ThemeService のコンストラクタ
func NewThemeService(r registry.Registry) *ThemeService {
	return &ThemeService{
		userRepo:  r.NewUserRepository(),
		themeRepo: r.NewThemeRepository(),
	}
}

// Create はお題を作成する
func (s *ThemeService) Create(ctx context.Context, db *gorp.DbMap, user model.User, sentence string) (model.Theme, error) {
	theme, err := s.themeRepo.Create(ctx, db, user, sentence)
	if err != nil {
		return model.Theme{}, err
	}
	return theme, nil
}

// List は自分が作成したお題を取得する
func (s *ThemeService) List(ctx context.Context, db *gorp.DbMap, user model.User) ([]model.Theme, error) {
	themes, err := s.themeRepo.FetchByUser(ctx, db, user)
	if err != nil {
		return []model.Theme{}, err
	}
	return themes, nil
}
