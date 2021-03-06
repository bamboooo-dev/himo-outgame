package registry

import (
	himo_mysql "github.com/bamboooo-dev/himo-outgame/internal/interface/mysql/himo"
	himo_repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/bamboooo-dev/himo-outgame/pkg/env"
	"go.uber.org/zap"
)

// Registry は DI コンテナ
type Registry interface {
	Config() *env.Config
	NewUserRepository() himo_repo.UserRepository
	NewThemeRepository() himo_repo.ThemeRepository
	NewHistoryRepository() himo_repo.HistoryRepository
}

type registry struct {
	config *env.Config
	l      *zap.SugaredLogger
}

// NewRegistry is Registry constructor.
func NewRegistry(cfg *env.Config, l *zap.SugaredLogger) Registry {
	return &registry{cfg, l}
}

func (r *registry) NewUserRepository() himo_repo.UserRepository {
	return himo_mysql.NewUserRepositoryMysql(r.l)
}

func (r *registry) NewThemeRepository() himo_repo.ThemeRepository {
	return himo_mysql.NewThemeRepositoryMysql(r.l)
}

func (r *registry) NewHistoryRepository() himo_repo.HistoryRepository {
	return himo_mysql.NewHistoryRepositoryMysql(r.l)
}

func (r *registry) Config() *env.Config {
	if r.config == nil {
		return &env.Config{}
	}
	return r.config
}
