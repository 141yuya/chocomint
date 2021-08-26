package main

import (
	"github.com/141yuya/go-clean-architecture/src/di"
	"github.com/141yuya/go-clean-architecture/src/domain/entities"
	"github.com/141yuya/go-clean-architecture/src/infrastructure"
)

func main() {
	sqlHandler := infrastructure.NewSqlHandler()
	sqlHandler.DB.AutoMigrate(entities.User{})
	router := infrastructure.NewRouter(di.InitUserController(sqlHandler))
	router.Start()
}
