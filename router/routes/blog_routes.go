package routes

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/aifuxi/fgo/internal/middleware"
	"github.com/aifuxi/fgo/internal/repository"
	"github.com/aifuxi/fgo/internal/service"
	"github.com/aifuxi/fgo/pkg/db"
	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(api *gin.RouterGroup, svc service.UserService) {
	h := handler.NewBlogHandler(service.NewBlogService(repository.NewBlogRepository(db.GetDB())))

	routes := api.Group("/blogs")
	routes.Use(middleware.Auth())
	{
		routes.GET("", middleware.RequirePermissions(svc, "blog::list"), h.List)
		routes.POST("", middleware.RequirePermissions(svc, "blog::create"), h.Create)

		routes.GET("/:id", middleware.RequirePermissions(svc, "blog::view"), h.FindByID)
		routes.PUT("/:id", middleware.RequirePermissions(svc, "blog::update"), h.UpdateByID)
		routes.DELETE("/:id", middleware.RequirePermissions(svc, "blog::delete"), h.DeleteByID)
	}
}
