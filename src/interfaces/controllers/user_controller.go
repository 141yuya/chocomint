package controllers

import (
	"net/http"
	"strconv"

	"github.com/141yuya/go-clean-architecture/domain/entities"
	"github.com/141yuya/go-clean-architecture/usecases"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (userController *UserController) Create(c *gin.Context) {
	u := entities.User{}
	c.Bind(&u)
	user, err := userController.UserUsecase.Add(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (userController *UserController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := userController.UserUsecase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (userController *UserController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := entities.User{}
	c.Bind(&u)

	user, err := userController.UserUsecase.Update(id, &u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (userController *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := userController.UserUsecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
