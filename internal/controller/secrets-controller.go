package controller

import (
	"net/http"
	"prival-api/internal/entity"
	"prival-api/internal/middleware"
	"prival-api/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SecretsController interface {
	CreateSecret(*gin.Context)
	ListSecrets(*gin.Context)
	DeleteSecret(*gin.Context)
}

type secretsController struct {
	secretsService service.SecretsService
}

func NewSecretsController(secretsService service.SecretsService) SecretsController {
	return &secretsController{
		secretsService: secretsService,
	}
}

func (c *secretsController) CreateSecret(ctx *gin.Context) {
	req := struct {
		Description string `json:"description" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, _ := ctx.Get(middleware.GetAuthKey())
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user id must be numeric",
		})
		return
	}

	secret := &entity.Secret{
		UserID:      userIDInt,
		Description: req.Description,
	}

	err = c.secretsService.CreateSecret(secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}

func (c *secretsController) ListSecrets(ctx *gin.Context) {
	userID, _ := ctx.Get(middleware.GetAuthKey())
	userIDInt, err := strconv.Atoi(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user id must be numeric",
		})
		return
	}

	secrets, err := c.secretsService.ListSecrets(userIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"secrets": secrets,
	})
}

func (c *secretsController) DeleteSecret(ctx *gin.Context) {
	secretID := ctx.Param("id")
	secretIDInt, err := strconv.Atoi(secretID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "param secret_id was missing or not formatted correctly",
		})
		return
	}

	userID, _ := ctx.Get(middleware.GetAuthKey())
	deleted, err := c.secretsService.DeleteUsersSecretByID(userID.(int), secretIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"deleted": deleted,
	})
}
