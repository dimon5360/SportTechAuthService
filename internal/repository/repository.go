package repository

import (
	"app/main/internal/dto"
)

type UserRepositoryInterface interface {
	Init() error
	Create(interface{}) (interface{}, error)
	Read(req *dto.LoginPostgresRequest) (*dto.LoginPostgresResponse, error)
	Update(interface{}) (interface{}, error)
	Delete(interface{}) error
}

type TokenRepositoryInterface interface {
	Init() error
	Create(interface{}) (interface{}, error)
	Read(interface{}) (interface{}, error)
	Update(interface{}) (interface{}, error)
	Delete(interface{}) error
}
