package repository

import (
	"app/main/internal/repository"
	"app/main/pkg/env"
	"database/sql"
	"fmt"
	"os"
)

const (
	postgresConfigPathKey = "POSTGRES_CONFIG"

	postgresHostKey          = "POSTGRES_HOST"
	postgresAdminKey         = "POSTGRES_ADMIN"
	postgresAdminPasswordKey = "POSTGRES_ADMIN_PASSWORD"
	postgresDatabaseKey      = "POSTGRES_DATABASE"

	sslModeDisableExpression = "?sslmode=disable"
)

type userRepository struct {
	cli *sql.DB
}

func New() repository.Interface {
	return &userRepository{}
}

func (r *userRepository) Init() error {

	if err := env.Load(os.Getenv(postgresConfigPathKey)); err != nil {
		return err
	}

	opt := fmt.Sprintf("postgresql://%s:%s@%s/%s%s",
		os.Getenv(postgresAdminKey),
		os.Getenv(postgresAdminPasswordKey),
		os.Getenv(postgresHostKey),
		os.Getenv(postgresDatabaseKey),
		sslModeDisableExpression,
	)

	if cli, err := sql.Open("postgres", opt); err == nil && cli != nil {
		r.cli = cli
		return nil
	}

	return fmt.Errorf("can't create postgres client")
}

func (r *userRepository) Create(interface{}) (interface{}, error) {
	return nil, nil
}

func (r *userRepository) Read(interface{}) (interface{}, error) {
	return nil, nil
}

func (r *userRepository) Update(interface{}) (interface{}, error) {
	return nil, nil
}

func (r *userRepository) Delete(interface{}) error {
	return nil
}
