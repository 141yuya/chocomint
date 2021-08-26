//go:build wireinject
// +build wireinject

package di

import (
	"github.com/141yuya/chocomint/src/infrastructure"
	"github.com/141yuya/chocomint/src/interfaces/controllers"
	"github.com/141yuya/chocomint/src/interfaces/gateways"
	"github.com/141yuya/chocomint/src/usecases"
	"github.com/google/wire"
)

var superSet = wire.NewSet(
	gateways.NewUserRepository,
	usecases.NewUserUsecase,
	controllers.NewUserController,
)

func InitUserController(handler *infrastructure.SqlHandler) controllers.UserController {
	wire.Build(superSet)
	return controllers.UserController{}
}
