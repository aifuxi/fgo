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

	userService := service.NewUserService(repository.NewUserRepository(db.GetDB()))
	userHandler := handler.NewUserHandler(userService)

	authRoutes := apiV1.Group("/auth")
	{
		authRoutes.POST("/register", userHandler.Register)
		authRoutes.POST("/login", userHandler.Login)
	}

	userRoutes := apiV1.Group("/users")
	userRoutes.Use(middleware.Auth())
	{
		userRoutes.GET("", middleware.RequirePermissions(userService, "user::list"), userHandler.List)
		userRoutes.GET("/:id", middleware.RequirePermissions(userService, "user::view"), userHandler.FindByID)
		userRoutes.PUT("/:id", middleware.RequirePermissions(userService, "user::update"), userHandler.Update)
		userRoutes.DELETE("/:id", middleware.RequirePermissions(userService, "user::delete"), userHandler.Delete)
	}

	roleHandler := handler.NewRoleHandler(service.NewRoleService(repository.NewRoleRepository(db.GetDB())))

	roleRoutes := apiV1.Group("/roles")
	roleRoutes.Use(middleware.Auth())
	{
		roleRoutes.GET("", middleware.RequirePermissions(userService, "role::list"), roleHandler.List)
		roleRoutes.POST("", middleware.RequirePermissions(userService, "role::create"), roleHandler.Create)
		roleRoutes.GET("/:id", middleware.RequirePermissions(userService, "role::view"), roleHandler.FindByID)
		roleRoutes.PUT("/:id", middleware.RequirePermissions(userService, "role::update"), roleHandler.Update)
		roleRoutes.DELETE("/:id", middleware.RequirePermissions(userService, "role::delete"), roleHandler.Delete)
	}

	tagHandler := handler.NewTagHandler(service.NewTagService(repository.NewTagRepository(db.GetDB())))

	tagRoutes := apiV1.Group("/tags")
	tagRoutes.Use(middleware.Auth())
	{
		tagRoutes.GET("", middleware.RequirePermissions(userService, "tag::list"), tagHandler.List)
		tagRoutes.GET("/:id", middleware.RequirePermissions(userService, "tag::view"), tagHandler.FindByID)

		tagRoutes.POST("", middleware.RequirePermissions(userService, "tag::create"), tagHandler.Create)
		tagRoutes.PUT("/:id", middleware.RequirePermissions(userService, "tag::update"), tagHandler.UpdateByID)
		tagRoutes.DELETE("/:id", middleware.RequirePermissions(userService, "tag::delete"), tagHandler.DeleteByID)
	}

	blogHandler := handler.NewBlogHandler(service.NewBlogService(repository.NewBlogRepository(db.GetDB())))

	blogRoutes := apiV1.Group("/blogs")
	blogRoutes.Use(middleware.Auth())
	{
		blogRoutes.GET("", middleware.RequirePermissions(userService, "blog::list"), blogHandler.List)
		blogRoutes.GET("/:id", middleware.RequirePermissions(userService, "blog::view"), blogHandler.FindByID)

		blogRoutes.POST("", middleware.RequirePermissions(userService, "blog::create"), blogHandler.Create)
		blogRoutes.PUT("/:id", middleware.RequirePermissions(userService, "blog::update"), blogHandler.UpdateByID)
		blogRoutes.DELETE("/:id", middleware.RequirePermissions(userService, "blog::delete"), blogHandler.DeleteByID)
	}

	categoryHandler := handler.NewCategoryHandler(service.NewCategoryService(repository.NewCategoryRepository(db.GetDB())))

	categoryRoutes := apiV1.Group("/categories")
	categoryRoutes.Use(middleware.Auth())
	{
		categoryRoutes.GET("", middleware.RequirePermissions(userService, "category::list"), categoryHandler.List)
		categoryRoutes.GET("/:id", middleware.RequirePermissions(userService, "category::view"), categoryHandler.FindByID)

		categoryRoutes.POST("", middleware.RequirePermissions(userService, "category::create"), categoryHandler.Create)
		categoryRoutes.PUT("/:id", middleware.RequirePermissions(userService, "category::update"), categoryHandler.UpdateByID)
		categoryRoutes.DELETE("/:id", middleware.RequirePermissions(userService, "category::delete"), categoryHandler.DeleteByID)
	}

	return router
}
