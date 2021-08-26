//go:build wireinject
// +build wireinject

package main

import (
	"github.com/141yuya/chocomint/src/infrastructure"
	"github.com/141yuya/chocomint/src/interfaces/controllers"
	"github.com/141yuya/chocomint/src/interfaces/gateways"
	"github.com/141yuya/chocomint/src/usecases"
	"github.com/google/wire"
)

func InitUserController() controllers.UserController {
	wire.Build(
		infrastructure.NewSqlHandler,
		gateways.NewUserRepository,
		usecases.NewUserUsecase,
		controllers.NewUserController,
	)
	return controllers.UserController{}
}
