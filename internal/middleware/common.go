package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m Middleware) AbortWithError(c *gin.Context, err error) {
	if m.model.IsErrNotFound(err) {
		c.String(http.StatusNotFound, "%s", err)
	} else {
		c.String(http.StatusInternalServerError, "%s", err)
	}

	_ = c.Error(err)
	c.Abort()
}

func (m Middleware) AbortWithErrorAndBadRequestCode(c *gin.Context, err error) {
	c.String(http.StatusBadRequest, "%s", err)
	_ = c.Error(err)
	c.Abort()
}
