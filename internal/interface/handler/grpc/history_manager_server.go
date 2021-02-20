package grpc

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	"github.com/bamboooo-dev/himo-outgame/internal/usecase/interactor"
	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
	"github.com/bamboooo-dev/himo-outgame/pkg/grpcmiddleware"
	"github.com/go-gorp/gorp"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
)

// HistoryManagerServer gRPC サーバーの実装
type HistoryManagerServer struct {
	logger *zap.SugaredLogger
	lister *interactor.ListHistoriesInteractor
	db     *gorp.DbMap

	pb.UnimplementedHistoryManagerServer
}

// NewHistoryManagerServer returns HistoryManagerServer
func NewHistoryManagerServer(l *zap.SugaredLogger, r registry.Registry, db *gorp.DbMap) HistoryManagerServer {
	return HistoryManagerServer{
		logger: l,
		lister: interactor.NewListHistoriesInteractor(r),
		db:     db,
	}
}

// List gets my histories
func (s HistoryManagerServer) List(ctx context.Context, req *pb.ListHistoryRequest) (*pb.ListHistoryResponse, error) {

	userID := ctx.Value(grpcmiddleware.StringKey).(string)

	histories, err := s.lister.Call(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}
	pbhistories := make([]*pb.History, len(histories))
	for i, history := range histories {
		pbCreatedAt, err := ptypes.TimestampProto(history.CreatedAt)
		if err != nil {
			pbCreatedAt = nil
		}
		pbWithUsers := make([]*pb.User, len(history.WithUsers))
		for i, withUser := range history.WithUsers {
			pbWithUsers[i] = &pb.User{
				Id:       uint32(withUser.ID),
				Nickname: withUser.Nickname,
			}
		}
		pbhistories[i] = &pb.History{
			Id:        uint32(history.ID),
			CreatedAt: pbCreatedAt,
			WithUsers: pbWithUsers,
		}
	}
	ret := pb.ListHistoryResponse{
		Histories: pbhistories,
	}
	return &ret, nil
}
