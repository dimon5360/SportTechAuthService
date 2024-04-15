package repository

import (
	"app/main/internal/dto/models"
	"app/main/internal/repository"
	"app/main/pkg/env"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

const (
	redisConfigPathKey = "REDIS_CONFIG"

	redisHostKey          = "REDIS_HOST"
	redisAdminPasswordKey = "REDIS_ADMIN_PASSWORD"
	redisDatabaseKey      = "REDIS_DATABASE"
)

type tokenRepository struct {
	cli *redis.Client
}

func New() repository.TokenRepositoryInterface {
	return &tokenRepository{}
}

func (r *tokenRepository) Init() error {

	if err := env.Load(os.Getenv(redisConfigPathKey)); err != nil {
		return err
	}

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
	return nil, fmt.Errorf("method isn't implemented")
}

func (r *tokenRepository) ValidateRefreshToken(
	req *models.RefreshTokenRequestModel,
) error {
	return fmt.Errorf("method isn't implemented")
}

func (r *tokenRepository) ValidateAccessToken(
	req *models.AccessTokenRequestModel,
) error {
	return fmt.Errorf("method isn't implemented")
}

func (r *tokenRepository) GenerateTokens(
	req *models.GenerateTokensRequestModel,
) (*models.GenerateTokensResponseModel, error) {
	return nil, fmt.Errorf("method isn't implemented")
}
