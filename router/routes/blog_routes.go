package routes

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/aifuxi/fgo/internal/middleware"
	"github.com/aifuxi/fgo/internal/model"
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
		routes.GET("", middleware.RequirePermissions(svc, model.PermissionBlogList), h.List)
		routes.POST("", middleware.RequirePermissions(svc, model.PermissionBlogCreate), h.Create)

		routes.GET("/:id", middleware.RequirePermissions(svc, model.PermissionBlogView), h.FindByID)
		routes.PUT("/:id", middleware.RequirePermissions(svc, model.PermissionBlogUpdate), h.UpdateByID)
		routes.DELETE("/:id", middleware.RequirePermissions(svc, model.PermissionBlogDelete), h.DeleteByID)
	}
}
