package repository

import (
	"app/main/internal/dto/models"
	"app/main/internal/repository"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

const (
	redisHostKey          = "REDIS_HOST"
	redisAdminPasswordKey = "REDIS_ADMIN_PASSWORD"
	redisDatabaseKey      = "REDIS_DATABASE"

	methodNotImplemented string = "methos isn't implemented"
)

type tokenRepository struct {
	cli *redis.Client
}

func New() repository.TokenRepositoryInterface {
	return &tokenRepository{}
}

func (r *tokenRepository) Init() error {

	opt := redis.Options{
		Addr:     os.Getenv(redisHostKey),
		Password: os.Getenv(redisAdminPasswordKey),
		DB:       0, // use default DB
	}

	if cli := redis.NewClient(&opt); cli != nil {
		r.cli = cli
		return nil
	}
	return fmt.Errorf("can't create redis client")
}

func (r *tokenRepository) RefreshTokens(
	req *models.RefreshTokenRequestModel,
) (*models.RefreshTokenResponseModel, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *tokenRepository) ValidateRefreshToken(
	req *models.RefreshTokenRequestModel,
) error {
	return fmt.Errorf(methodNotImplemented)
}

func (r *tokenRepository) ValidateAccessToken(
	req *models.AccessTokenRequestModel,
) error {
	return fmt.Errorf(methodNotImplemented)
}

func (r *tokenRepository) GenerateTokens(
	req *models.GenerateTokensRequestModel,
) (*models.GenerateTokensResponseModel, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}
