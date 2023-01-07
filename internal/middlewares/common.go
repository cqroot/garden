package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, err error) {
	c.String(http.StatusInternalServerError, "%s", err)
	c.Error(err)
	c.Abort()
}
