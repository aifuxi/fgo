package router

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/aifuxi/fgo/internal/repository"
	"github.com/aifuxi/fgo/internal/service"
	"github.com/aifuxi/fgo/pkg/db"
	"github.com/aifuxi/fgo/pkg/response"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		response.Success(c, nil)
	})

	apiV1 := router.Group("/api/v1")

	tagHandler := handler.NewTagHandler(service.NewTagService(repository.NewTagRepository(db.GetDB())))

	tagRoutes := apiV1.Group("/tags")
	{
		tagRoutes.GET("", tagHandler.List)
		tagRoutes.POST("", tagHandler.Create)
		tagRoutes.GET("/:id", tagHandler.FindByID)
		tagRoutes.DELETE("/:id", tagHandler.DeleteByID)
		tagRoutes.PUT("/:id", tagHandler.UpdateByID)
	}

	return router
}
