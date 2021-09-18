package grpc

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	"github.com/bamboooo-dev/himo-outgame/internal/usecase/interactor"
	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
	"github.com/bamboooo-dev/himo-outgame/pkg/grpcmiddleware"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UserManagerServer gRPC サーバーの実装
type UserManagerServer struct {
	logger   *zap.SugaredLogger
	register *interactor.RegisterUserInteractor
	updater  *interactor.UpdateUserInteractor
	db       *gorp.DbMap

	pb.UnimplementedUserManagerServer
}

// NewUserManagerServer returns UserManagerServer
func NewUserManagerServer(l *zap.SugaredLogger, r registry.Registry, db *gorp.DbMap) UserManagerServer {
	return UserManagerServer{
		logger:   l,
		register: interactor.NewRegistUserInteractor(r),
		updater:  interactor.NewUpdateUserInteractor(r),
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
	// create and send header
	header := metadata.Pairs("access-token", accessToken)
	grpc.SendHeader(ctx, header)
	ret := pb.SignUpResponse{}
	return &ret, nil
}

func (s UserManagerServer) UpdateUserName(ctx context.Context, req *pb.UpdateUserNameRequest) (*pb.UpdateUserNameResponse, error) {
	nickname := req.GetNickname()
	userID := ctx.Value(grpcmiddleware.StringKey).(string)
	err := s.updater.Call(ctx, s.db, nickname, userID)
	if err != nil {
		return nil, err
	}
	ret := pb.UpdateUserNameResponse{}
	return &ret, nil

}

// AuthFuncOverride is to handle authentication
func (s UserManagerServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if fullMethodName == "/himo.v1.UserManager/SignUp" {
		return ctx, nil
	}

	ctx, err := grpcmiddleware.Authenticate(ctx)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}
