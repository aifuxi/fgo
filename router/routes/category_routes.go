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

func RegisterCategoryRoutes(api *gin.RouterGroup, svc service.UserService) {
	h := handler.NewCategoryHandler(service.NewCategoryService(repository.NewCategoryRepository(db.GetDB())))

	routes := api.Group("/categories")
	routes.Use(middleware.Auth())
	{
		routes.GET("", middleware.RequirePermissions(svc, model.PermissionCategoryList), h.List)
		routes.GET("/:id", middleware.RequirePermissions(svc, model.PermissionCategoryView), h.FindByID)

		routes.POST("", middleware.RequirePermissions(svc, model.PermissionCategoryCreate), h.Create)
		routes.PUT("/:id", middleware.RequirePermissions(svc, model.PermissionCategoryUpdate), h.UpdateByID)
		routes.DELETE("/:id", middleware.RequirePermissions(svc, model.PermissionCategoryDelete), h.DeleteByID)
	}
}
