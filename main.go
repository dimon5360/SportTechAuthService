package main

import (
	"app/main/proto"
	"app/main/utils"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/// TODO:
/// 1. handle request with dbms
/// 2. transfer response

type AuthUsersService struct {
	proto.UnimplementedAuthUsersServiceServer
}

func main() {

	env := utils.Init()
	env.Load("env/app.env")
	env.Load("env/db.env")

	fmt.Println("Auth service v." + env.Value("VERSION_APP"))

	// storage := storage.CreateStorage()
	// storage.Init(env.Value("CONNECTION_STRING") + "?sslmode=disable")

	// user := storage.GetUserById("1")
	// fmt.Println(user)

	lis, err := net.Listen("tcp", "localhost:40402")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption

	serv := AuthUsersService{}
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterAuthUsersServiceServer(grpcServer, &serv)
	grpcServer.Serve(lis)
}

func (s *AuthUsersService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {

	return &proto.User{
		Id:        "1",
		Username:  "default",
		Password:  "default",
		Email:     "default",
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}, nil
}
