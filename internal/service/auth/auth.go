package service

import (
	"app/main/internal/service"
	"context"
	"fmt"
	"net"
	"os"

	proto "proto/go"

	"google.golang.org/grpc"
)

const (
	serviceHostKey = "SERVICE_HOST"
)

type authService struct {
	grpc *grpcService
}

type grpcService struct {
	proto.UnimplementedAuthServiceServer
}

func New() service.Interface {
	return &authService{}
}

func (s *authService) Init() error {

	host := os.Getenv(serviceHostKey)
	if len(host) == 0 {
		return fmt.Errorf("service host not found")
	}

	lis, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	service := &grpcService{}

	proto.RegisterAuthServiceServer(grpcServer, service)

	s.grpc = service
	return grpcServer.Serve(lis)
}

func (s *authService) Run() error {
	return nil
}

func (s *grpcService) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	return nil, nil
}

func (s *grpcService) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	return nil, nil
}

func (s *grpcService) RefreshToken(ctx context.Context, req *proto.RefreshTokenRequest) (*proto.RefreshTokenResponse, error) {
	return nil, nil
}

func (s *grpcService) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	return nil, nil
}
