package repository

import (
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

func (r *tokenRepository) Create(interface{}) (interface{}, error) {
	return nil, nil
}

func (r *tokenRepository) Read(interface{}) (interface{}, error) {
	return nil, nil
}

func (r *tokenRepository) Update(interface{}) (interface{}, error) {
	return nil, nil
}

func (r *tokenRepository) Delete(interface{}) error {
	return nil
}
