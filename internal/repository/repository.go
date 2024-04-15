package repository

import (
	"app/main/internal/dto/models"
)

type UserRepositoryInterface interface {
	Init() error
	Create(req *models.RegisterPostgresRequestModel) (*models.RegisterPostgresResponseModel, error)
	Read(req *models.LoginPostgresRequestModel) (*models.LoginPostgresResponseModel, error)
	Update(interface{}) (interface{}, error)
	Delete(interface{}) error
}

type TokenRepositoryInterface interface {
	Init() error
	RefreshTokens(req *models.RefreshTokenRequestModel) (*models.RefreshTokenResponseModel, error)
	ValidateRefreshToken(req *models.RefreshTokenRequestModel) error
	ValidateAccessToken(req *models.AccessTokenRequestModel) error
	GenerateTokens(req *models.GenerateTokensRequestModel) (*models.GenerateTokensResponseModel, error)
}
