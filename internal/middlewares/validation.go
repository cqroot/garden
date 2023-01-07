package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateIdMiddleware(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		AbortWithError(c, err)
		return
	}

	c.Set("id", id)
}
