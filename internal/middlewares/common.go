package middlewares

import (
	"net/http"

	"github.com/cqroot/garden/internal/models"
	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, err error) {
	if models.IsErrNotFound(err) {
		c.String(http.StatusNotFound, "%s", err)
	} else {
		c.String(http.StatusInternalServerError, "%s", err)
	}

	_ = c.Error(err)
	c.Abort()
}

func AbortWithErrorAndBadRequestCode(c *gin.Context, err error) {
	c.String(http.StatusBadRequest, "%s", err)
	_ = c.Error(err)
	c.Abort()
}
