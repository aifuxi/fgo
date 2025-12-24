package model

import (
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        int64          `gorm:"primarykey" json:"id,string"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
