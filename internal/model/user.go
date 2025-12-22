package model

type User struct {
	CommonModel
	Nickname string  `gorm:"size:100;not null;comment:用户昵称" json:"nickname"`
	Email    string  `gorm:"size:100;uniqueIndex;not null;comment:邮箱" json:"email"`
	Password string  `gorm:"size:255;not null;comment:密码" json:"-"`
	Roles    []*Role `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}
