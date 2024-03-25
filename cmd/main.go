package main

import (
	"app/main/internal/storage"
	"app/main/internal/utils"
	"log"
	"net"

	proto "github.com/dimon5360/SportTechProtos/gen/go/proto"
	"google.golang.org/grpc"
)

const (
	configPath  = "../SportTechDockerConfig/"
	serviceEnv  = ".env"
	postgresEnv = configPath + "postgres.env"
	redisEnv    = configPath + "redis.env"
)

func main() {

	env := utils.Env()
	env.Load(serviceEnv, postgresEnv, redisEnv)

	log.Println("SportTech auth service v." + env.Value("SERVICE_VERSION"))

	service := storage.CreateService()
	service.Init()

	lis, err := net.Listen("tcp", env.Value("AUTH_GRPC_HOST"))
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterAuthUsersServiceServer(grpcServer, service)
	grpcServer.Serve(lis)
}
