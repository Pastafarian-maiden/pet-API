package internal

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.Use(authMiddleware())
	router.Run(os.Getenv("PORT_API"))
}
