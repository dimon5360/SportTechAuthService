package repository

import (
	"app/main/internal/dto"
	"app/main/internal/repository"
	"app/main/pkg/env"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	postgresConfigPathKey string = "POSTGRES_CONFIG"

	postgresHostKey          string = "POSTGRES_HOST"
	postgresAdminKey         string = "POSTGRES_USER"
	postgresAdminPasswordKey string = "POSTGRES_PASSWORD"
	postgresDatabaseKey      string = "POSTGRES_DB"

	sslModeDisableExpression string = "?sslmode=disable"

	methodNotImplemented string = "methos isn't implemented"
)

type userRepository struct {
	cli *sql.DB
}

func New() repository.UserRepositoryInterface {
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

	log.Println(opt)

	cli, err := sql.Open("postgres", opt)
	if err == nil && cli != nil {
		r.cli = cli
		log.Println("postgres connection is initialized")
		return nil
	}

	return err
}

func (r *userRepository) Create(interface{}) (interface{}, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) Read(req *dto.LoginPostgresRequest) (*dto.LoginPostgresResponse, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) Update(interface{}) (interface{}, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) Delete(interface{}) error {
	return fmt.Errorf(methodNotImplemented)
}
