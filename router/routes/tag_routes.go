package routes

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/aifuxi/fgo/internal/middleware"
	"github.com/aifuxi/fgo/internal/repository"
	"github.com/aifuxi/fgo/internal/service"
	"github.com/aifuxi/fgo/pkg/db"
	"github.com/gin-gonic/gin"
)

func RegisterTagRoutes(api *gin.RouterGroup, svc service.UserService) {
	h := handler.NewTagHandler(service.NewTagService(repository.NewTagRepository(db.GetDB())))

	routes := api.Group("/tags")
	routes.Use(middleware.Auth())
	{
		routes.GET("", middleware.RequirePermissions(svc, "tag::list"), h.List)
		routes.POST("", middleware.RequirePermissions(svc, "tag::create"), h.Create)

		routes.GET("/:id", middleware.RequirePermissions(svc, "tag::view"), h.FindByID)
		routes.PUT("/:id", middleware.RequirePermissions(svc, "tag::update"), h.UpdateByID)
		routes.DELETE("/:id", middleware.RequirePermissions(svc, "tag::delete"), h.DeleteByID)
	}
}
