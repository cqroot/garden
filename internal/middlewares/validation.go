package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func IdValidationMiddleware(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		AbortWithErrorAndBadRequestCode(c, err)
		return
	}

	c.Set("id", uint(id))
}
