package app

import (
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

	version := os.Getenv(serviceVersionKey)
	if len(version) == 0 {
		return nil, fmt.Errorf("service version not found")
	}

	fmt.Println("SportTech auth service v." + version)
	log.Println("provider initialized")

	return p.initAuthService()

}

func (p *provider) initAuthService() (service.Interface, error) {

	service := authService.New()

	if err := service.Init(); err != nil {
		return nil, err
	}

	log.Println("auth service created")
	return service, nil
}
