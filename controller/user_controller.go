package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masa720/todo-backend-golang/model"
	"github.com/masa720/todo-backend-golang/usecase"
)

type UserController interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu}
}

func (uc *userController) Signup(c *gin.Context) {
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userUsecase.Signup(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (uc *userController) Signin(c *gin.Context) {
	var loginData struct {
		Mail string `json:"mail" binding:"required"`
		Pass string `json:"pass" binding:"required"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.userUsecase.Signin(loginData.Mail, loginData.Pass)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
