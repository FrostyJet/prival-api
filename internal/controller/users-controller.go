package controller

import (
	"fmt"
	"net/http"
	"prival-api/internal/entity"
	"prival-api/internal/middleware"
	"prival-api/internal/service"
	"prival-api/pkg/token"
	"time"

	"github.com/gin-gonic/gin"
)

type UsersController interface {
	ListUsers(*gin.Context)
	RegisterUser(*gin.Context)
	LoginUser(*gin.Context)
}

type usersController struct {
	usersService service.UsersService
	token        token.Token
}

func NewUsersController(usersService service.UsersService, token token.Token) UsersController {
	return &usersController{
		usersService: usersService,
		token:        token,
	}
}

func (c *usersController) ListUsers(ctx *gin.Context) {
	userID, _ := ctx.Get(middleware.GetAuthKey())

	ctx.JSON(http.StatusOK, gin.H{
		"users":        c.usersService.GetUsers(),
		"current-user": userID,
	})
}

func (c *usersController) RegisterUser(ctx *gin.Context) {
	user := &entity.User{}

	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.usersService.CreateUser(user)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (c *usersController) LoginUser(ctx *gin.Context) {
	req := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password are required",
		})
		return
	}

	user, err := c.usersService.LoginUser(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := c.token.Create(fmt.Sprint(user.ID), time.Minute*30)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}
