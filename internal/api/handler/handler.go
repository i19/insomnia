package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"insomnia/pkg/errors"
)

type ResponseProtocol struct {
	Code    errors.ErrorCode // business code, 0 means ok
	Message string           // error message
	Data    interface{}      // result
}

func Response(c *gin.Context, code errors.ErrorCode, errorMessage string, data interface{}) {
	c.JSON(http.StatusOK, ResponseProtocol{
		Code:    code,
		Message: errorMessage,
		Data:    data,
	})
}
func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseProtocol{
		Code:    errors.OK,
		Message: "success",
		Data:    data,
	})
}
