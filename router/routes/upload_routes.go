package routes

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUploadRoutes(api *gin.RouterGroup) {
	h := handler.NewUploadHandler()

	authRoutes := api.Group("/upload")
	{
		authRoutes.POST("/presign", h.UploadPresign)
	}
}
