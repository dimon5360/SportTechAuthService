package models

import (
	"time"
)

const (
	UserRole string = "User"
)

type LoginPostgresRequest struct {
	Email string
	Role  string
}

type LoginPostgresResponse struct {
	Id          uint64
	Email       string
	Password    string
	Role        string
	ProfileId   uint64
	IsValidated bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RegisterPostgresRequest struct {
	Email    string
	Password string
	Role     string
}

type RegisterPostgresResponse struct {
	Id          uint64
	Email       string
	Password    string
	Role        string
	ProfileId   uint64
	IsValidated bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
