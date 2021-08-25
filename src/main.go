package main

import (
	"github.com/141yuya/chocomint/src/infrastructure"
)

func main() {
	userController := InitController()
	infrastructure.InitRouter(userController)
}
