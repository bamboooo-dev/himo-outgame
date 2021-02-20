package interactor

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

type GetContentInteractor struct {
	userRepo himo_repo.UserRepository
}

func NewGetContentInteractor(r registry.Registry) *GetContentInteractor {
	return &GetContentInteractor{
		userRepo: r.NewUserRepository(),
	}
}

func (g *GetContentInteractor) Call(ctx context.Context, db *gorp.DbMap, userID string) error {
	_, err := g.userRepo.Find(ctx, db, userID)
	if err != nil {
		return err
	}
	return nil
}
