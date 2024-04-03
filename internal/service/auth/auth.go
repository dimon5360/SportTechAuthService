package service

import (
	"app/main/internal/repository"
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
	// grpc            *grpcService
	grpc            *grpc.Server
	userReposoitory repository.Interface
	tokenRepository repository.Interface
}

type grpcService struct {
	proto.UnimplementedAuthServiceServer
}

func New(userReposoitory repository.Interface, tokenRepository repository.Interface) service.Interface {
	return &authService{
		grpc:            nil,
		userReposoitory: userReposoitory,
		tokenRepository: tokenRepository,
	}
}

func (s *authService) Init() error {

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	proto.RegisterAuthServiceServer(grpcServer, &grpcService{})

	s.grpc = grpcServer
	return nil
}

func (s *authService) Run() error {

	lis, err := net.Listen("tcp", os.Getenv(serviceHostKey))
	if err != nil {
		return err
	}

	return s.grpc.Serve(lis)
}

func (s *grpcService) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {

	fmt.Println(req)

	// 1. validate user credentials from postgres
	// 2. generate tokens and store in redis

	return nil, nil
}

func (s *grpcService) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {

	// 1. create new user in postgres
	return nil, nil
}

func (s *grpcService) RefreshToken(ctx context.Context, req *proto.RefreshTokenRequest) (*proto.RefreshTokenResponse, error) {

	// 1. validate user refresh token from redis
	// 2. generate new tokens and store in redis
	return nil, nil
}

func (s *grpcService) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {

	// 1. read user access token from redis
	// 2. check token is valid
	return nil, nil
}
