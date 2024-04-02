package main

import (
	"app/main/internal/app"
	"log"
)

const (
	configPath  = "../SportTechDockerConfig/"
	serviceEnv  = ".env"
	postgresEnv = configPath + "postgres.env"
	redisEnv    = configPath + "redis.env"
)

func main() {

	a := app.New()

	if err := a.Init(); err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
