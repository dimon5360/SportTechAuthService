package tdo

import "time"

type PostgresUser struct {
	Id         uint64
	Email      string
	Password   string
	Role       string
	Created_at time.Time
	Updated_at time.Time
}
