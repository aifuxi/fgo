package router

import (
	"github.com/aifuxi/fgo/internal/handler"
	"github.com/aifuxi/fgo/internal/middleware"
	"github.com/aifuxi/fgo/internal/repository"
	"github.com/aifuxi/fgo/internal/service"
	"github.com/aifuxi/fgo/pkg/db"
	"github.com/aifuxi/fgo/pkg/response"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		response.Success(c, gin.H{
			"status": "ok",
			// TODO 从配置文件中读取版本号
			"version": "v0.0.1",
		})
	})

	apiV1 := router.Group("/api/v1")

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db.GetDB())))

	authRoutes := apiV1.Group("/auth")
	{
		authRoutes.POST("/register", userHandler.Register)
		authRoutes.POST("/login", userHandler.Login)
	}

	userRoutes := apiV1.Group("/users")
	userRoutes.Use(middleware.Auth())
	{
		userRoutes.GET("", userHandler.List)
		userRoutes.GET("/:id", userHandler.FindByID)
		userRoutes.PUT("/:id", userHandler.Update)
		userRoutes.DELETE("/:id", userHandler.Delete)
	}

	roleHandler := handler.NewRoleHandler(service.NewRoleService(repository.NewRoleRepository(db.GetDB())))

	roleRoutes := apiV1.Group("/roles")
	roleRoutes.Use(middleware.Auth())
	{
		roleRoutes.GET("", roleHandler.List)
		roleRoutes.POST("", roleHandler.Create)
		roleRoutes.GET("/:id", roleHandler.FindByID)
		roleRoutes.PUT("/:id", roleHandler.Update)
		roleRoutes.DELETE("/:id", roleHandler.Delete)
	}

	tagHandler := handler.NewTagHandler(service.NewTagService(repository.NewTagRepository(db.GetDB())))

	tagRoutes := apiV1.Group("/tags")
	tagRoutes.Use(middleware.Auth())
	{
		tagRoutes.GET("", tagHandler.List)
		tagRoutes.POST("", tagHandler.Create)
		tagRoutes.GET("/:id", tagHandler.FindByID)
		tagRoutes.DELETE("/:id", tagHandler.DeleteByID)
		tagRoutes.PUT("/:id", tagHandler.UpdateByID)
	}

	blogHandler := handler.NewBlogHandler(service.NewBlogService(repository.NewBlogRepository(db.GetDB())))

	blogRoutes := apiV1.Group("/blogs")
	blogRoutes.Use(middleware.Auth())
	{
		blogRoutes.GET("", blogHandler.List)
		blogRoutes.POST("", blogHandler.Create)
		blogRoutes.GET("/:id", blogHandler.FindByID)
		blogRoutes.DELETE("/:id", blogHandler.DeleteByID)
		blogRoutes.PUT("/:id", blogHandler.UpdateByID)
	}

	categoryHandler := handler.NewCategoryHandler(service.NewCategoryService(repository.NewCategoryRepository(db.GetDB())))

	categoryRoutes := apiV1.Group("/categories")
	categoryRoutes.Use(middleware.Auth())
	{
		categoryRoutes.GET("", categoryHandler.List)
		categoryRoutes.POST("", categoryHandler.Create)
		categoryRoutes.GET("/:id", categoryHandler.FindByID)
		categoryRoutes.DELETE("/:id", categoryHandler.DeleteByID)
		categoryRoutes.PUT("/:id", categoryHandler.UpdateByID)
	}

	return router
}
