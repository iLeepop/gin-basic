package bootstrap

import (
	"gin-basic/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(c *Container) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	router := app.Group("/api")

	router.Use(middleware.Trace())
	router.Use(middleware.AccessLog())
	router.Use(middleware.Cors())

	registerServerRoutes(router, c)

	return app
}

func registerServerRoutes(router *gin.RouterGroup, c *Container) {
	router.GET("/health", c.ServerController.HealthCheck)
}
