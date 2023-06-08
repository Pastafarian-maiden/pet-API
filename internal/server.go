package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	router := gin.Default()
	router.Use(authMiddleware())

	api := router.Group("/api")

	router.GET("/health", GetHealth)
	api.GET("/names", GetNames)
	api.POST("/names", PostNames)
	api.GET("/addresses", GetAddresses)
	api.POST("/addresses", PostAdresses)

	router.Run(viper.GetString("ADDR_API"))
}
