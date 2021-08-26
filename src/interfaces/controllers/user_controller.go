package controllers

import (
	"strconv"

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

func (userController *UserController) Index(c *gin.Context) {
	users, err := userController.UserUsecase.GetUsers()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (userController *UserController) Create(c *gin.Context) {
	u := entities.User{}
	c.Bind(&u)
	err := userController.UserUsecase.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, nil)
}

func (userController *UserController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := userController.UserUsecase.GetUser(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}

func (userController *UserController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := entities.User{}
	c.Bind(&u)

	user, err := userController.UserUsecase.Update(id, u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}
