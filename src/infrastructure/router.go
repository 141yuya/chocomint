package infrastructure

import (
	"github.com/141yuya/chocomint/src/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	UserController controllers.UserController
}

func (r Router) Start() {
	engine := gin.Default()
	u := engine.Group("/users")
	{
		u.GET("", r.UserController.Index)
		u.GET("/:id", r.UserController.Show)
		u.POST("", r.UserController.Create)
		// u.PUT("/:id", ctrl.Update)
		// u.DELETE("/:id", ctrl.Delete)
	}
	engine.Run()
}

func NewRouter(userController controllers.UserController) Router {
	return Router{UserController: userController}
}
