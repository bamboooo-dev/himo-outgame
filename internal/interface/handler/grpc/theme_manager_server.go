package grpc

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	"github.com/bamboooo-dev/himo-outgame/internal/usecase/interactor"
	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
	"github.com/bamboooo-dev/himo-outgame/pkg/grpcmiddleware"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// ThemeManagerServer gRPC サーバーの実装
type ThemeManagerServer struct {
	logger  *zap.SugaredLogger
	creator *interactor.CreateThemeInteractor
	db      *gorp.DbMap

	pb.UnimplementedThemeManagerServer
}

// NewThemeManagerServer returns ThemeManagerServer
func NewThemeManagerServer(l *zap.SugaredLogger, r registry.Registry, db *gorp.DbMap) ThemeManagerServer {
	return ThemeManagerServer{
		logger:  l,
		creator: interactor.NewCreateThemeInteractor(r),
		db:      db,
	}
}

// Create creates theme
func (s ThemeManagerServer) Create(ctx context.Context, req *pb.ThemeRequest) (*pb.ThemeResponse, error) {

	sentence := req.GetSentence()
	userID := ctx.Value(grpcmiddleware.StringKey).(string)

	theme, err := s.creator.Call(ctx, s.db, userID, sentence)
	if err != nil {
		return nil, err
	}
	ret := pb.ThemeResponse{
		Theme: &pb.Theme{
			Id:       uint32(theme.ID),
			Sentence: theme.Sentence,
		},
	}
	return &ret, nil
}
