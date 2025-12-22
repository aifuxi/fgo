package routes

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/aifuxi/fgo/internal/middleware"
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
		routes.GET("", middleware.RequirePermissions(svc, "role::list"), h.List)
		routes.POST("", middleware.RequirePermissions(svc, "role::create"), h.Create)

		routes.GET("/:id", middleware.RequirePermissions(svc, "role::view"), h.FindByID)
		routes.PUT("/:id", middleware.RequirePermissions(svc, "role::update"), h.Update)
		routes.DELETE("/:id", middleware.RequirePermissions(svc, "role::delete"), h.Delete)
	}
}
