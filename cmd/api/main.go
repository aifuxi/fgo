package main

import (
	"net/http"

	"github.com/aifuxi/fgo/internal/model"
	"github.com/aifuxi/fgo/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitMySQL()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/test", func(c *gin.Context) {
		tag := model.Tag{
			Name: "test",
			Slug: "test",
		}

		if err := db.DB.Create(&tag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    tag,
		})
	})

	router.Run("127.0.0.1:8080") // 默认监听 0.0.0.0:8080
}
