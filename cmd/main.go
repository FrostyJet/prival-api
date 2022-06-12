package main

import (
	"prival-api/internal/http"
	"prival-api/internal/service"

	"github.com/gin-gonic/gin"
)

var ginRouter http.Router

func main() {
	service.InitDB()

	engine := gin.Default()

	ginRouter = http.NewGinRouter(engine)
	ginRouter.InitRoutes()
	ginRouter.Serve()
}
