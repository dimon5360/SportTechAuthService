package main

import (
	"app/main/proto"
	"app/main/storage"
	"app/main/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

/// TODO:
/// 1. handle request with dbms
/// 2. transfer response

func main() {

	env := utils.Init()
	env.Load("env/app.env")
	env.Load("env/db.env")

	log.Println("Auth service v." + env.Value("VERSION_APP"))

	service := storage.CreateService()
	service.Init(env.Value("CONNECTION_STRING") + "?sslmode=disable")

	lis, err := net.Listen("tcp", "localhost:40402")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterAuthUsersServiceServer(grpcServer, &service)
	grpcServer.Serve(lis)
}
