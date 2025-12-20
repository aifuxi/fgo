package main

import (
	"github.com/aifuxi/fgo/internal/model"
	"github.com/aifuxi/fgo/pkg/db"
)

func main() {
	db.InitMySQL()

	db.DB.AutoMigrate(&model.Tag{})
}
