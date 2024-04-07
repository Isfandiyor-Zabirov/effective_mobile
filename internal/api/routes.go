package api

import (
	"effective_mobile_tech_task/internal/api/handlers"
	"effective_mobile_tech_task/logger"
	"effective_mobile_tech_task/utils/env"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RunApp launches all routes
// @title Technical task
// @version 1.0
// @description Open API документация для проекта VISOR
// @host localhost
// @BasePath /
func RunApp() {
	settings := env.GetSettings()
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())
	router.MaxMultipartMemory = 8 << 20
	logger.FormatLogs(router)

	v1 := router.Group("/api/v1/")
	v1.GET("ping", ping)

	{
		carRoutes := v1.Group("/cars")
		carRoutes.POST("", handlers.CreateCars)
		carRoutes.PUT("", handlers.UpdateCar)
		carRoutes.DELETE("/:id", handlers.DeleteCar)
		carRoutes.GET("", handlers.GetCars)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"reason": "No route found"})
	})

	err := router.Run(settings.AppParams.Port)
	if err != nil {
		logger.Error.Fatal("Cannot start service:", err.Error())
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"reason": "Welcome to Project"})
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, auth")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}
