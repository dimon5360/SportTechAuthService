package repository

import (
	"app/main/internal/dto/models"
)

type UserRepositoryInterface interface {
	Init() error
	Create(req *models.RegisterPostgresRequest) (*models.RegisterPostgresResponse, error)
	Read(req *models.LoginPostgresRequest) (*models.LoginPostgresResponse, error)
	Update(interface{}) (interface{}, error)
	Delete(interface{}) error
}

type TokenRepositoryInterface interface {
	Init() error
	Refresh(interface{}) (interface{}, error)
	Validate(interface{}) (interface{}, error)
}
