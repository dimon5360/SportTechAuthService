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

const methodNotImplemented string = "method not implemented"

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

func (s *GrpcEndpoint) LoginUser(
	ctx context.Context, req *proto.LoginUserRequest,
) (*proto.LoginUserResponse, error) {

	log.Println("user login procedure")
	if s.userReposoitory == nil {
		log.Fatal("user repository isn't initialized")
	}

	user, err := s.userReposoitory.GetUser(
		models.ConvertRequestLoginModel(req, models.UserRole))
	if err != nil {
		return nil, err
	}

	salt := os.Getenv("PASSWORD_HASH_SALT")
	hash, err := HashPassword(req.Password, salt)
	if err != nil {
		return nil, err
	}

	if isValid := ValidatePassword(user.Password, hash); !isValid {
		return nil, fmt.Errorf("invalid creadentials")
	}

	tokens, err := s.tokenRepository.GenerateTokens(
		&models.GenerateTokensRequestModel{
			UserId: user.Id,
		})

	if err != nil {
		return nil, err
	}
	return models.ConvertResponseLoginModel(user, tokens), nil
}

func (s *GrpcEndpoint) RegisterUser(
	ctx context.Context, req *proto.RegisterUserRequest,
) (*proto.RegisterUserResponse, error) {

	log.Println("user register request")
	if s.userReposoitory == nil {
		log.Fatal("user repository isn't initialized")
	}

	salt := os.Getenv("PASSWORD_HASH_SALT")
	_, err := HashPassword(req.Password, salt)
	if err != nil {
		return nil, fmt.Errorf("hashing password failed")
	}

	_, err = s.userReposoitory.CreateUser(
		models.ConvertRequestRegisterModel(req, models.UserRole))
	if err != nil {
		return nil, err
	}

	return models.ConvertResponseRegisterModel("success"), nil
}

func (s *GrpcEndpoint) RefreshToken(
	ctx context.Context, req *proto.RefreshTokenRequest,
) (*proto.RefreshTokenResponse, error) {

	log.Println("refresh token request")

	model := models.ConvertRequestRefreshTokenModel(req)
	err := s.tokenRepository.ValidateRefreshToken(model)
	if err != nil {
		return nil, err
	}

	response, err := s.tokenRepository.RefreshTokens(model)
	if err != nil {
		return nil, err
	}

	return models.ConvertResponseRefreshTokenModel(response), nil
}

func (s *GrpcEndpoint) ValidateToken(
	ctx context.Context, req *proto.ValidateTokenRequest,
) (*proto.ValidateTokenResponse, error) {

	log.Println("access token validating request")

	model := models.ConvertRequestValidateTokenModel(req)
	err := s.tokenRepository.ValidateAccessToken(model)
	return models.ConvertResponseValidateTokenModel(err), err
}

func (s *GrpcEndpoint) CreateProfile(
	context.Context, *proto.CreateProfileRequest,
) (*proto.ProfileResponse, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (s *GrpcEndpoint) DeleteProfile(
	context.Context, *proto.DeleteProfileRequest,
) (*proto.ProfileResponse, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (s *GrpcEndpoint) GetProfile(
	context.Context, *proto.GetProfileRequest,
) (*proto.ProfileResponse, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (s *GrpcEndpoint) UpdateProfile(
	context.Context, *proto.UpdateProfileRequest,
) (*proto.ProfileResponse, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}
