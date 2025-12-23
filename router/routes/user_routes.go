package routes

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/aifuxi/fgo/internal/middleware"
	"github.com/aifuxi/fgo/internal/model"
	"github.com/aifuxi/fgo/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup, svc service.UserService) {
	h := handler.NewUserHandler(svc)

	routes := api.Group("/users")
	routes.Use(middleware.Auth())
	{
		routes.GET("", middleware.RequirePermissions(svc, model.PermissionAdminAll, model.PermissionUserList), h.List)

		routes.GET("/:id", middleware.RequirePermissions(svc, model.PermissionAdminAll, model.PermissionUserView), h.FindByID)
		routes.PUT("/:id", middleware.RequirePermissions(svc, model.PermissionAdminAll, model.PermissionUserUpdate), h.Update)
		routes.DELETE("/:id", middleware.RequirePermissions(svc, model.PermissionAdminAll, model.PermissionUserDelete), h.Delete)
	}
}
