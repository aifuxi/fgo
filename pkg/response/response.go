package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Success 成功
func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

// ParamError 校验参数错误
func ParamError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    1001,
		Message: msg,
		Data:    nil,
	})
}

// BusinessError 业务逻辑错误
func BusinessError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    -1,
		Message: msg,
		Data:    nil,
	})
}
