package main

import (
	"context"
	"log"
	"net"
	userApi "project/internal/api/user"
	"project/internal/config"
	userRepo "project/internal/repository/user"
	userServ "project/internal/service/user"
	desc "project/pkg/user_v1"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	configPath = "config/config.yaml"
)

func main() {
	ctx := context.Background()
	mainConfig, err := config.MainConfigInit(configPath)
	if err != nil {
		log.Fatalf("failed loading config:%s", err)
	}

	pool, err := pgxpool.New(ctx, mainConfig.DbConfigLoad())
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping to database: %s", err)
	}

	log.Printf("database connection is successfull")

	repo := userRepo.NewRepository(pool)
	serv := userServ.NewService(repo)
	api := userApi.NewUserImplementation(serv)

	lis, err := net.Listen("tcp", mainConfig.ServerConfigLoader())
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, api)

	log.Printf("server is listening at %s", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
