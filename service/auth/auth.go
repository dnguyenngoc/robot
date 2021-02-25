package auth

import (

	"errors"
	
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dnguyenngoc/robot/service/exceptions"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			exceptions.CustomError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}