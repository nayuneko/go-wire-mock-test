// +build wireinject

package main

// The build tag makes sure the stub is not built in the final build.
import (
	"github.com/google/wire"
	"github.com/nayuneko/go-wire-mock-test/repository"
	"github.com/nayuneko/go-wire-mock-test/service"
)

func InitializeService() service.UserService {
	wire.Build(service.NewUserService, repository.NewUserRepository, repository.NewEntryRepository)
	return nil
}
