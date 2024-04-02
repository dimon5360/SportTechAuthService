package models

import (
	"time"
)

type User struct {
	Id         uint64
	Email      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
}
