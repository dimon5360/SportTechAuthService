package app

import (
	"app/main/internal/endpoint"
	tokenRepository "app/main/internal/repository/token"
	userRepository "app/main/internal/repository/user"
	"app/main/internal/service"
	authService "app/main/internal/service/auth"
	"fmt"
	"log"
	"os"

	"app/main/pkg/env"
)

type ProviderInterface interface {
	Init() (service.Interface, error)
}

type provider struct {
}

func NewServiceProvider() ProviderInterface {
	return &provider{}
}

func (p *provider) Init() (service.Interface, error) {

	if err := env.Init(); err != nil {
		return nil, err
	}

	fmt.Println("SportTech user service v." + os.Getenv(serviceVersionKey))
	log.Println("provider initialized")

	return p.initAuthService()
}

func (p *provider) initAuthService() (service.Interface, error) {

	userRepo := userRepository.New()
	if err := userRepo.Init(); err != nil {
		log.Fatal(err.Error())
	}

	tokenRepo := tokenRepository.New()
	if err := tokenRepo.Init(); err != nil {
		log.Fatal(err.Error())
	}

	service := authService.New(endpoint.New(userRepo, tokenRepo))

	if err := service.Init(); err != nil {
		return nil, err
	}

	log.Println("user service created")
	return service, nil
}
