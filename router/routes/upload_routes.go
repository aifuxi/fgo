package routes

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/aifuxi/fgo/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterUploadRoutes(api *gin.RouterGroup) {
	h := handler.NewUploadHandler(service.NewUploadService())

	authRoutes := api.Group("/upload")
	{
		authRoutes.POST("/presign", h.UploadPresign)
	}
}
