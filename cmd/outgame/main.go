package main

import (
	"context"
	"fmt"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	handler "github.com/bamboooo-dev/himo-outgame/internal/interface/handler/grpc"
	"github.com/bamboooo-dev/himo-outgame/internal/interface/mysql"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	"github.com/bamboooo-dev/himo-outgame/pkg/env"
	pb "github.com/bamboooo-dev/himo-outgame/pkg/grpc/v1/himo/proto"
	"github.com/bamboooo-dev/himo-outgame/pkg/grpcmiddleware"
)

// from LDFLAGS
var revision = "undefined"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic '%v' captured\n", err)
		}
		// graceful shutdown をするために待つ処理
	}()

	fmt.Printf("Version is %s\n", revision)

	cfg, err := env.LoadConfigFromTemplate()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	himoDB, err := mysql.NewDB(cfg.HimoMySQL)
	if err != nil {
		sugar.Error(ctx, err)
		return
	}
	defer func() {
		if err := himoDB.Db.Close(); err != nil {
			sugar.Error(ctx, err)
			return
		}
	}()

	lis, err := net.Listen("tcp", "0.0.0.0:5502")
	if err != nil {
		wrapErr := fmt.Errorf("failed to listen: %w", err)
		sugar.Error(wrapErr)
		return
	}
	sugar.Info("started to listen gRPC :5502")

	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(grpcmiddleware.Authenticate)))

	reflection.Register(server)

	registry := registry.NewRegistry(cfg, sugar)

	pb.RegisterRoomServer(server, handler.NewRoomServer(registry, himoDB))
	pb.RegisterUserManagerServer(server, handler.NewUserManagerServer(sugar, registry, himoDB))
	pb.RegisterThemeManagerServer(server, handler.NewThemeManagerServer(sugar, registry, himoDB))
	pb.RegisterHistoryManagerServer(server, handler.NewHistoryManagerServer(sugar, registry, himoDB))
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
