package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping to check
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
