package exceptions

import "github.com/gin-gonic/gin"

func CustomError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code: status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}