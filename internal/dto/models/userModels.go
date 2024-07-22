package models

import (
	proto "proto/go"
	"time"
)

const (
	UserRole string = "User"
)

type LoginPostgresRequestModel struct {
	Email string
	Role  string
}

type LoginPostgresResponseModel struct {
	Id          uint64
	Email       string
	Password    string
	Role        string
	IsValidated bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ConvertRequestLoginModel(
	req *proto.LoginUserRequest,
	role string,
) *LoginPostgresRequestModel {
	return &LoginPostgresRequestModel{
		Email: req.Email,
		Role:  role,
	}
}

func ConvertResponseLoginModel(
	user *LoginPostgresResponseModel,
	tokens *GenerateTokensResponseModel,
) *proto.LoginUserResponse {
	return &proto.LoginUserResponse{
		Id:           user.Id,
		AccessToken:  ConvertTokensModel(&tokens.AccessToken),
		RefreshToken: ConvertTokensModel(&tokens.RefreshToken),
		IsValidated:  user.IsValidated,
		Error:        "success",
	}
}

type RegisterPostgresRequestModel struct {
	Email    string
	Password string
	Role     string
}

type RegisterPostgresResponseModel struct {
	Id          uint64
	Email       string
	Password    string
	Role        string
	ProfileId   uint64
	IsValidated bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ConvertRequestRegisterModel(
	req *proto.RegisterUserRequest,
	role string,
) *RegisterPostgresRequestModel {
	return &RegisterPostgresRequestModel{
		Email:    req.Email,
		Password: req.Password,
		Role:     role,
	}
}

func ConvertResponseRegisterModel(result string) *proto.RegisterUserResponse {
	return &proto.RegisterUserResponse{
		Error: result,
	}
}
