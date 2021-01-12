package main

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	handler "github.com/bamboooo-dev/himo-outgame/internal/interface/handler/grpc"
	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	lis, err := net.Listen("tcp", "0.0.0.0:5502")
	if err != nil {
		wrapErr := fmt.Errorf("failed to listen: %w", err)
		sugar.Error(wrapErr)
		return
	}

	server := grpc.NewServer()

	reflection.Register(server)

	pb.RegisterRoomServer(server, handler.NewRoomServer())
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
