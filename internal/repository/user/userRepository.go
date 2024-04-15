package repository

import (
	"app/main/internal/dto/models"
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

func (r *userRepository) Create(
	req *models.RegisterPostgresRequestModel,
) (*models.RegisterPostgresResponseModel, error) {

	template := "INSERT into users (email, password, role, is_validated)" +
		"values ('%s', '%s', '%s', %t) ON conflict DO NOTHING RETURNING *;"
	sql := fmt.Sprintf(template, req.Email, req.Password, req.Role, false)

	rows := r.cli.QueryRow(sql)
	if err := rows.Err(); err != nil {
		log.Print(err)
		return &models.RegisterPostgresResponseModel{}, err
	}

	var model models.RegisterPostgresResponseModel
	// err := rows.Scan(&model.Id, &model.Email, &model.Password, &model.Role, &model.ProfileId, &model.IsValidated, &model.CreatedAt, &model.UpdatedAt)
	err := rows.Scan(&model)
	if err != nil {
		log.Print(err)
		return &models.RegisterPostgresResponseModel{}, err
	}
	return &model, nil
}

func (r *userRepository) Read(
	req *models.LoginPostgresRequestModel,
) (*models.LoginPostgresResponseModel, error) {

	template := "SELECT * users WHERE email=%s AND role=%s;"
	sql := fmt.Sprintf(template, req.Email, req.Role)

	rows := r.cli.QueryRow(sql)
	if err := rows.Err(); err != nil {
		log.Print(err)
		return &models.LoginPostgresResponseModel{}, err
	}

	var model models.LoginPostgresResponseModel
	err := rows.Scan(&model)
	if err != nil {
		log.Print(err)
		return &models.LoginPostgresResponseModel{}, err
	}
	return &model, nil
}

func (r *userRepository) Update(interface{}) (interface{}, error) {
	return nil, fmt.Errorf(methodNotImplemented)
}

func (r *userRepository) Delete(interface{}) error {
	return fmt.Errorf(methodNotImplemented)
}
