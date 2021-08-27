//go:build wireinject
// +build wireinject

package di

import (
	"github.com/141yuya/go-clean-architecture/domain/repositories"
	"github.com/141yuya/go-clean-architecture/infrastructure"
	"github.com/141yuya/go-clean-architecture/interfaces/controllers"
	"github.com/141yuya/go-clean-architecture/interfaces/gateways"
	"github.com/141yuya/go-clean-architecture/usecases"
	"github.com/google/wire"
)

var superSet = wire.NewSet(
	gateways.NewUserRepository,
	wire.Bind(new(repositories.UserRepository), new(*gateways.UserRepository)),
	usecases.NewUserUsecase,
	controllers.NewUserController,
)

func InitUserController(handler *infrastructure.SqlHandler) controllers.UserController {
	wire.Build(superSet)
	return controllers.UserController{}
}
