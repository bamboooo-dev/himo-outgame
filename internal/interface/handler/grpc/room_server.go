package grpc

import (
	"context"

	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
)

// RoomServer gRPC サーバーの実装
type RoomServer struct {
	pb.UnimplementedRoomServer
}

// NewRoomServer returns RoomServer
func NewRoomServer() RoomServer {
	return RoomServer{}
}

// GetContent gets room content
func (s RoomServer) GetContent(ctx context.Context, req *pb.ContentRequest) (*pb.ContentResponse, error) {
	ret := pb.ContentResponse{
		Table: &pb.Table{
			Id:   uint64(23),
			Name: "HIKARI table",
		},
	}
	return &ret, nil
}
