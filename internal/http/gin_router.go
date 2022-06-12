package http

import (
	"prival-api/internal/controller"
	"prival-api/internal/middleware"
	"prival-api/internal/service"
	"prival-api/pkg/token"

	"github.com/gin-gonic/gin"
)

type ginRouter struct {
	engine            *gin.Engine
	token             token.Token
	usersController   controller.UsersController
	secretsController controller.SecretsController
}

func NewGinRouter(engine *gin.Engine) Router {
	usersService := service.NewUsersService(service.DB)
	secretsService := service.NewSecretsService(service.DB)

	jwtToken := token.NewJWTToken("secret")

	return &ginRouter{
		engine:            engine,
		token:             jwtToken,
		usersController:   controller.NewUsersController(usersService, jwtToken),
		secretsController: controller.NewSecretsController(secretsService),
	}
}

func (r *ginRouter) InitRoutes() {
	router := r.engine
	authRequired := router.Group("/").Use(middleware.Auth(r.token))

	// Users
	authRequired.GET("/users", r.usersController.ListUsers)
	router.POST("/users/register", r.usersController.RegisterUser)
	router.POST("/users/login", r.usersController.LoginUser)

	// Secrets
	authRequired.GET("/secrets", r.secretsController.ListSecrets)
	authRequired.POST("/secrets", r.secretsController.CreateSecret)
	authRequired.DELETE("/secrets/:id", r.secretsController.DeleteSecret)
}

func (r *ginRouter) Serve() {
	r.engine.Run()
}
