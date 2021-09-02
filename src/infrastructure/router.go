package infrastructure

import (
	"github.com/141yuya/go-clean-architecture/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	UserController controllers.UserController
}

func NewRouter(userController controllers.UserController) Router {
	return Router{UserController: userController}
}

func (r *Router) SetupRouter() *gin.Engine {
	engine := gin.Default()
	u := engine.Group("/users")
	{
		u.GET("", r.UserController.Index)
		u.GET("/:id", r.UserController.Show)
		u.POST("", r.UserController.Create)
		u.PUT("/:id", r.UserController.Update)
		u.DELETE("/:id", r.UserController.Delete)
	}
	return engine
}
