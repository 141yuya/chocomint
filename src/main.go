package main

import (
	"github.com/141yuya/chocomint/src/di"
	"github.com/141yuya/chocomint/src/domain/entities"
	"github.com/141yuya/chocomint/src/infrastructure"
)

func main() {
	sqlHandler := infrastructure.NewSqlHandler()
	sqlHandler.DB.AutoMigrate(entities.User{})
	router := infrastructure.NewRouter(di.InitUserController(sqlHandler))
	router.Start()
}
