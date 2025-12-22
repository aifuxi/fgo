package model

type Permission struct {
	CommonModel
	Name        string  `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Code        string  `gorm:"size:100;uniqueIndex;not null" json:"code"`
	Description string  `gorm:"size:255" json:"description"`
	Roles       []*Role `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}
