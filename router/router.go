package router

import (
	"github.com/aifuxi/fgo/internal/repository"
	"github.com/aifuxi/fgo/internal/service"
	"github.com/aifuxi/fgo/pkg/db"
	"github.com/aifuxi/fgo/pkg/response"
	"github.com/aifuxi/fgo/router/routes"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		response.Success(c, gin.H{
			"status": "ok",
			// TODO 从配置文件中读取版本号
			"version": "0.0.1",
		})
	})

	apiV1 := router.Group("/api/v1")

	userService := service.NewUserService(repository.NewUserRepository(db.GetDB()))

	routes.RegisterAuthRoutes(apiV1, userService)
	routes.RegisterUserRoutes(apiV1, userService)
	routes.RegisterRoleRoutes(apiV1, userService)
	routes.RegisterTagRoutes(apiV1, userService)
	routes.RegisterBlogRoutes(apiV1, userService)
	routes.RegisterCategoryRoutes(apiV1, userService)

	return router
}
