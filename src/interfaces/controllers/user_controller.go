package controllers

import (
	"github.com/141yuya/chocomint/src/domain/entities"
	"github.com/141yuya/chocomint/src/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecases.UserUsecase
}

func NewUserController(u usecases.UserUsecase) UserController {
	return UserController{UserUsecase: u}
}

func (userController *UserController) Create(c *gin.Context) {
	u := entities.User{}
	c.Bind(&u)
	err := userController.UserUsecase.Add(u)
	if err != nil {
		c.JSON(500, NewError("", err))
		return
	}
	c.JSON(201, nil)
}
