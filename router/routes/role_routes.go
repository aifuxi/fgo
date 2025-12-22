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

func RegisterRoleRoutes(api *gin.RouterGroup, svc service.UserService) {
	h := handler.NewRoleHandler(service.NewRoleService(repository.NewRoleRepository(db.GetDB())))

	routes := api.Group("/roles")
	routes.Use(middleware.Auth())
	{
		routes.GET("", middleware.RequirePermissions(svc, model.PermissionRoleList), h.List)
		routes.POST("", middleware.RequirePermissions(svc, model.PermissionRoleCreate), h.Create)

		routes.GET("/:id", middleware.RequirePermissions(svc, model.PermissionRoleView), h.FindByID)
		routes.PUT("/:id", middleware.RequirePermissions(svc, model.PermissionRoleUpdate), h.Update)
		routes.DELETE("/:id", middleware.RequirePermissions(svc, model.PermissionRoleDelete), h.Delete)
	}
}
