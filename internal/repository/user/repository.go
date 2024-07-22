package repository

import (
	"app/main/internal/dto/models"
	"app/main/internal/repository"
	"app/main/pkg/env"
	"context"
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

	QUERY_ROW_ERROR  string = "Query row error: %v"
	SCAN_MODEL_ERROR string = "Scan model error: %v"
)

type userRepository struct {
	cli *sql.DB
}

func New() repository.UserRepositoryInterface {
	return &userRepository{}
}

func testRepositoryConnection(cli *sql.DB) {

	rows := cli.QueryRowContext(context.Background(), "select version();")
	if err := rows.Err(); err != nil {
		log.Fatalf(QUERY_ROW_ERROR, err)
		return
	}

	var version string
	if err := rows.Scan(&version); err != nil {
		log.Fatalf(SCAN_MODEL_ERROR, err)
		return
	}

	log.Printf("version database is %s", version)
	log.Println("Connection is established")
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
		testRepositoryConnection(cli)
		log.Println("postgres connection is stable")
		return nil
	}
	return err
}

func (r *userRepository) CreateUser(
	req *models.RegisterPostgresRequestModel,
) (*models.RegisterPostgresResponseModel, error) {

	hash, err := repository.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("hashing password failed")
	}

	template := "INSERT into users (email, password, role, is_validated) " +
		"values ('%s', '%s', '%s', %t) ON conflict DO NOTHING;"
	sql := fmt.Sprintf(template, req.Email, hash, req.Role, false)

	rows := r.cli.QueryRowContext(context.Background(), sql)
	if err := rows.Err(); err != nil {
		log.Printf(QUERY_ROW_ERROR, err)
		return nil, err
	}

	return &models.RegisterPostgresResponseModel{}, nil
}

func (r *userRepository) GetUser(
	req *models.LoginPostgresRequestModel,
) (*models.LoginPostgresResponseModel, error) {

	template := "SELECT * from users WHERE email='%s' AND role='%s';"
	sql := fmt.Sprintf(template, req.Email, req.Role)

	rows := r.cli.QueryRowContext(context.Background(), sql)
	if err := rows.Err(); err != nil {
		log.Printf(QUERY_ROW_ERROR, err)
		return nil, err
	}

	var model models.LoginPostgresResponseModel
	err := rows.Scan(
		&model.Id,
		&model.Email,
		&model.Password,
		&model.Role,
		&model.IsValidated,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if err != nil {
		log.Printf(SCAN_MODEL_ERROR, err)
		return nil, err
	}
	return &model, nil
}

func (r *userRepository) UpdateUser(interface{}) (interface{}, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) DeleteUser(interface{}) error {
	return fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) CreateProfile(interface{}) (interface{}, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) GetProfile(interface{}) (interface{}, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) UpdateProfile(interface{}) (interface{}, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) DeleteProfile(interface{}) error {
	return fmt.Errorf(methodNotImplemented)
}
