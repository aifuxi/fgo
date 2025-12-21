package main

import (
	"github.com/aifuxi/fgo/pkg/db"
	"github.com/aifuxi/fgo/router"
)

func main() {

	db.InitMySQL()

	router := router.Setup()

	router.Run("127.0.0.1:8080") // 默认监听 0.0.0.0:8080
}
