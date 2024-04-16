package repository

import (
	"app/main/internal/dto/models"
)

type UserRepositoryInterface interface {
	Init() error
	CreateUser(req *models.RegisterPostgresRequestModel) (*models.RegisterPostgresResponseModel, error)
	GetUser(req *models.LoginPostgresRequestModel) (*models.LoginPostgresResponseModel, error)
	UpdateUser(interface{}) (interface{}, error)
	DeleteUser(interface{}) error

	CreateProfile(interface{}) (interface{}, error)
	GetProfile(interface{}) (interface{}, error)
	UpdateProfile(interface{}) (interface{}, error)
	DeleteProfile(interface{}) error
}

type TokenRepositoryInterface interface {
	Init() error
	RefreshTokens(req *models.RefreshTokenRequestModel) (*models.RefreshTokenResponseModel, error)
	ValidateRefreshToken(req *models.RefreshTokenRequestModel) error
	ValidateAccessToken(req *models.AccessTokenRequestModel) error
	GenerateTokens(req *models.GenerateTokensRequestModel) (*models.GenerateTokensResponseModel, error)
}
