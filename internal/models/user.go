package models

import (
	"log"
	"time"

	"github.com/dimon5360/SportTechProtos/gen/go/proto"
	"golang.org/x/crypto/bcrypt"
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
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println(err)
	}
	return err == nil
}