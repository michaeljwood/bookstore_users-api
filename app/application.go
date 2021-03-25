package app

import (
	"github.com/gin-gonic/gin"
	"github.com/michaeljwood/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("About to start the application")

	router.Run(":8080")
}
