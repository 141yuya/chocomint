package main

import (
	"github.com/141yuya/go-clean-architecture/di"
	"github.com/141yuya/go-clean-architecture/domain/entities"
	"github.com/141yuya/go-clean-architecture/infrastructure"
)

func main() {
	sqlHandler := infrastructure.NewSqlHandler()
	sqlHandler.DB.AutoMigrate(entities.User{})
	router := infrastructure.NewRouter(di.InitUserController(sqlHandler))
	engine := router.SetupRouter()
	engine.Run()
}
