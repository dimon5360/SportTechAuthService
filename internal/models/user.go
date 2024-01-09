package models

import (
	"time"

	"github.com/dimon5360/SportTechProtos/gen/go/proto"
)

type User struct {
	Id         uint64
	Username   string
	Password   string
	Email      string
	Created_at time.Time
	Updated_at time.Time
}


func (user* User) ValidateCredentials(req *proto.AuthUserRequest) bool {
	return user.Email == req.Email && user.Password == req.Password
}