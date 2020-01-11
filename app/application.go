package app

import (
	"github.com/fdiaz7/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication entry point
func StartApplication() {
	mapUrls()
	logger.Info("about to start application")
	router.Run(":8080")
}
