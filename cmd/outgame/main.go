package main

import (
	"log"

	"google.golang.org/grpc"

	handler "github.com/bamboooo-dev/himo-outgame/internal/interface/handler/grpc"
	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
)

func main() {
	server := grpc.NewServer()
	pb.RegisterRoomServer(server, handler.NewRoomServer())
	log.Fatal("test")
}
