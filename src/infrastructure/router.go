package infrastructure

import (
	"github.com/141yuya/chocomint/src/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter(userController controllers.UserController) {
	r := router(userController)
	r.Run()
}

func router(userController controllers.UserController) *gin.Engine {
	router := gin.Default()
	u := router.Group("/users")
	{
		// u.GET("", ctrl.Index)
		// u.GET("/:id", ctrl.Show)
		u.POST("", userController.Create)
		// u.PUT("/:id", ctrl.Update)
		// u.DELETE("/:id", ctrl.Delete)
	}
	return router
}
