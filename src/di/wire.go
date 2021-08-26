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

func InitUserController(handler *infrastructure.SqlHandler) controllers.UserController {
	wire.Build(
		gateways.NewUserRepository,
		usecases.NewUserUsecase,
		controllers.NewUserController,
	)
	return controllers.UserController{}
}
