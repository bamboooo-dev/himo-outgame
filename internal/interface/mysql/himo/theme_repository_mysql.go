package himo

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	dao "github.com/bamboooo-dev/himo-outgame/internal/interface/dao/himo"
	repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// ThemeRepositoryMysql は ThemeRepository の MySQL 実装
type ThemeRepositoryMysql struct {
	logger *zap.SugaredLogger
}

// NewThemeRepositoryMysql は ThemeRepositoryMysql のコンストラクタ
func NewThemeRepositoryMysql(l *zap.SugaredLogger) repo.ThemeRepository {
	return ThemeRepositoryMysql{logger: l}
}

// Create inserts a theme
func (t ThemeRepositoryMysql) Create(ctx context.Context, db *gorp.DbMap, user model.User, sentence string) (model.Theme, error) {
	themeDAO := &dao.Theme{
		Sentence: sentence,
		UserID:   user.ID,
	}

	err := db.Insert(themeDAO)
	if err != nil {
		return model.Theme{}, err
	}

	theme := model.Theme{
		ID:       themeDAO.ID,
		Sentence: themeDAO.Sentence,
		Creator:  user,
	}
	return theme, nil
}
