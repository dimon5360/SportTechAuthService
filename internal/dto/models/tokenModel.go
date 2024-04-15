package models

import (
	proto "proto/go"
)

type TokenModel struct {
	value string
	age   uint64
}

type RefreshTokenRequestModel struct {
	UserId       uint64
	RefreshToken string
}

type RefreshTokenResponseModel struct {
	UserId       uint64
	AccessToken  TokenModel
	RefreshToken TokenModel
	Error        error
}

func ConvertRequestRefreshTokenModel(
	req *proto.RefreshTokenRequest,
) *RefreshTokenRequestModel {
	return &RefreshTokenRequestModel{
		UserId:       req.Id,
		RefreshToken: req.RefreshToken,
	}
}

func ConvertResponseRefreshTokenModel(
	resp *RefreshTokenResponseModel,
) *proto.RefreshTokenResponse {
	return &proto.RefreshTokenResponse{
		Id: resp.UserId,
		AccessToken: &proto.Token{
			Value: resp.AccessToken.value,
			Age:   resp.AccessToken.age,
		},
		RefreshToken: &proto.Token{
			Value: resp.RefreshToken.value,
			Age:   resp.RefreshToken.age,
		},
		Error: resp.Error.Error(),
	}
}

type AccessTokenRequestModel struct {
	UserId      uint64
	AccessToken string
}
type AccessTokenResponseModel struct {
	IsValid bool
	Error   error
}

func ConvertRequestValidateTokenModel(
	req *proto.ValidateTokenRequest,
) *AccessTokenRequestModel {
	return &AccessTokenRequestModel{
		UserId:      req.Id,
		AccessToken: req.AccessToken,
	}
}

func ConvertResponseValidateTokenModel(
	err error,
) *proto.ValidateTokenResponse {
	return &proto.ValidateTokenResponse{
		IsValid: err == nil,
		Error:   err.Error(),
	}
}

type GenerateTokensRequestModel struct {
	UserId uint64
}

type GenerateTokensResponseModel struct {
	AccessToken  TokenModel
	RefreshToken TokenModel
}

func ConvertTokensModel(token *TokenModel) *proto.Token {
	return &proto.Token{
		Value: token.value,
		Age:   token.age,
	}
}
