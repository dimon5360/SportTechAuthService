package service

import (
	"app/main/internal/endpoint"
	"app/main/internal/service"
	"net"
	"os"

	proto "proto/go"

	"google.golang.org/grpc"
)

const (
	serviceHostKey = "SERVICE_HOST"
)

type authService struct {
	grpcServer   *grpc.Server
	userEndpoint *endpoint.GrpcEndpoint
}

func New(userEndpoint *endpoint.GrpcEndpoint) service.Interface {
	return &authService{
		userEndpoint: userEndpoint,
	}
}

func (s *authService) Init() error {

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterAuthServiceServer(grpcServer, s.userEndpoint)

	s.grpcServer = grpcServer
	return nil
}

func (s *authService) Run() error {

	lis, err := net.Listen("tcp", os.Getenv(serviceHostKey))
	if err != nil {
		return err
	}

	return s.grpcServer.Serve(lis)
}
