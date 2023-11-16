package models

import "time"

type User struct {
	Id         uint64
	Username   string
	Password   string
	Email      string
	Created_at time.Time
	Updated_at time.Time
}
