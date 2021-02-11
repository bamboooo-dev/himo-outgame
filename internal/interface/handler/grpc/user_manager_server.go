package grpc

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	"github.com/bamboooo-dev/himo-outgame/internal/usecase/interactor"
	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// UserManagerServer gRPC サーバーの実装
type UserManagerServer struct {
	logger   *zap.SugaredLogger
	register *interactor.RegisterUserInteractor
	db       *gorp.DbMap

	pb.UnimplementedUserManagerServer
}

// NewUserManagerServer returns UserManagerServer
func NewUserManagerServer(l *zap.SugaredLogger, r registry.Registry, db *gorp.DbMap) UserManagerServer {
	return UserManagerServer{
		logger:   l,
		register: interactor.NewRegistUserInteractor(r),
		db:       db,
	}
}

// SignUp registers user
func (s UserManagerServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	nickName := req.GetNickname()
	accessToken, err := s.register.Call(ctx, s.db, nickName)
	if err != nil {
		return nil, err
	}
	ret := pb.SignUpResponse{
		AccessToken: accessToken,
	}
	return &ret, nil
}
