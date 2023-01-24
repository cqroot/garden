package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (m Middleware) IdValidationMiddleware(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		m.AbortWithErrorAndBadRequestCode(c, err)
		return
	}

	c.Set("id", uint(id))
}
