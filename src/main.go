package main

import (
	"github.com/141yuya/chocomint/src/domain/entities"
	"github.com/141yuya/chocomint/src/infrastructure"
	"github.com/141yuya/chocomint/src/interfaces/controllers"
	"github.com/141yuya/chocomint/src/interfaces/gateways"
	"github.com/141yuya/chocomint/src/usecases"
)

func main() {
	sqlHandler := infrastructure.NewSqlHandler()
	sqlHandler.DB.AutoMigrate(entities.User{})
	router := infrastructure.NewRouter(newUserController(sqlHandler))
	router.Start()
}

func newUserController(handler *infrastructure.SqlHandler) controllers.UserController {
	userRepository := gateways.NewUserRepository(handler)
	userUsecase := usecases.NewUserUsecase(userRepository)
	return controllers.NewUserController(userUsecase)
}
