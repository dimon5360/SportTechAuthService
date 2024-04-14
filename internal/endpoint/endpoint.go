package endpoint

import (
	"app/main/internal/dto/models"
	"app/main/internal/repository"
	"context"
	"fmt"
	"log"
	"os"
	proto "proto/go"
)

type GrpcEndpoint struct {
	proto.UnimplementedAuthServiceServer

	userReposoitory repository.UserRepositoryInterface
	tokenRepository repository.TokenRepositoryInterface
}

func New(userReposoitory repository.UserRepositoryInterface,
	tokenRepository repository.TokenRepositoryInterface,
) *GrpcEndpoint {
	return &GrpcEndpoint{
		userReposoitory: userReposoitory,
		tokenRepository: tokenRepository,
	}
}

func (s *GrpcEndpoint) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {

	log.Println("user login procedure")
	if s.userReposoitory == nil {
		log.Fatal("user repository isn't initialized")
	}

	response, err := s.userReposoitory.Read(
		&models.LoginPostgresRequest{
			Email: req.Email,
			Role:  models.UserRole,
		})

	if err != nil {
		return nil, err
	}

	salt := os.Getenv("PASSWORD_HASH_SALT")
	hash, err := HashPassword(req.Password, salt)
	if err != nil {
		return &proto.LoginUserResponse{
			Error: proto.AuthError_INVALID_CREDENTIALS,
		}, nil
	}

	if isValid := ValidatePassword(response.Password, hash); !isValid {
		return &proto.LoginUserResponse{
			Error: proto.AuthError_INVALID_CREDENTIALS,
		}, nil
	}

	return &proto.LoginUserResponse{
		Id:           1,
		AccessToken:  &proto.Token{},
		RefreshToken: &proto.Token{},
		ProfileId:    1,
		IsValidated:  false,
		Error:        proto.AuthError_OK,
	}, nil
}

func (s *GrpcEndpoint) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {

	log.Println("user register request")
	if s.userReposoitory == nil {
		log.Fatal("user repository isn't initialized")
	}

	salt := os.Getenv("PASSWORD_HASH_SALT")
	_, err := HashPassword(req.Password, salt)
	if err != nil {
		return nil, fmt.Errorf("hashing password failed")
	}

	_, err = s.userReposoitory.Create(
		&models.RegisterPostgresRequest{
			Email:    req.Email,
			Password: req.Password,
			Role:     models.UserRole,
		})

	if err != nil {
		return &proto.RegisterUserResponse{
			Error: proto.AuthError_NOT_FOUND,
		}, err
	}

	return &proto.RegisterUserResponse{
		Error: proto.AuthError_OK,
	}, nil
}

func (s *GrpcEndpoint) RefreshToken(ctx context.Context, req *proto.RefreshTokenRequest) (*proto.RefreshTokenResponse, error) {

	// 1. validate user refresh token from redis
	// 2. generate new tokens and store in redis
	return nil, nil
}

func (s *GrpcEndpoint) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {

	// 1. read user access token from redis
	// 2. check token is valid
	return nil, nil
}
