package http

import (
	"log"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mrfade/case-sss/internal/adapters/configs"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	configManager *configs.ConfigManager,
) (*Router, error) {
	if configManager.IsProduction() {
		log.Println("INFO: Running router in production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	ginConfig := cors.DefaultConfig()
	allowedOrigins := configManager.Container.HTTP.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve() error {
	return r.Run()
}
