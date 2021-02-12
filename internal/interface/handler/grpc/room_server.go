package grpc

import (
	"context"

	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	"github.com/bamboooo-dev/himo-outgame/internal/usecase/interactor"
	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
	"github.com/bamboooo-dev/himo-outgame/pkg/grpcmiddleware"
	"github.com/go-gorp/gorp"
)

// RoomServer gRPC サーバーの実装
type RoomServer struct {
	db     *gorp.DbMap
	getter *interactor.GetContentInteractor
	pb.UnimplementedRoomServer
}

// NewRoomServer returns RoomServer
func NewRoomServer(r registry.Registry, db *gorp.DbMap) RoomServer {
	return RoomServer{
		getter: interactor.NewGetContentInteractor(r),
		db:     db,
	}
}

// GetContent gets room content
func (s RoomServer) GetContent(ctx context.Context, req *pb.ContentRequest) (*pb.ContentResponse, error) {
	userID := ctx.Value(grpcmiddleware.StringKey).(string)
	err := s.getter.Call(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}
	ret := pb.ContentResponse{
		Table: &pb.Table{
			Id:   uint64(23),
			Name: "HIKARI table",
		},
	}
	return &ret, nil
}
