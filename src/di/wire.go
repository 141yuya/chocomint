//go:build wireinject
// +build wireinject

package di

import (
	"github.com/141yuya/chocomint/src/domain/repositories"
	"github.com/141yuya/chocomint/src/infrastructure"
	"github.com/141yuya/chocomint/src/interfaces/controllers"
	"github.com/141yuya/chocomint/src/interfaces/gateways"
	"github.com/141yuya/chocomint/src/usecases"
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
