package model

type Tag struct {
	CommonModel

	Name string `gorm:"size:255;unique;not null;comment:标签名称" json:"name"`
	Slug string `gorm:"size:255;uniqueIndex;not null;comment:标签 slug" json:"slug"`

	Blogs []*Blog `gorm:"many2many:blog_tags;" json:"blogs,omitempty"`
}
