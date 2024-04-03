package env

import (
	gt "github.com/joho/godotenv"
)

func Init() error {
	return gt.Load()
}

func Load(filename ...string) error {
	return gt.Load(filename...)
}
