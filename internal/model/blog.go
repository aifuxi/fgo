package model

type Blog struct {
	CommonModel

	Title       string `gorm:"size:255;unique;not null;comment:博客标题" json:"title"`
	Slug        string `gorm:"size:255;uniqueIndex;not null;comment:博客 slug" json:"slug"`
	Description string `gorm:"not null;comment:博客描述" json:"description"`
	Cover       string `gorm:"comment:博客封面" json:"cover"`
	Content     string `gorm:"not null;comment:博客内容" json:"content"`

	CategoryID uint      `gorm:"comment:分类ID" json:"categoryID"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`

	Tags []*Tag `gorm:"many2many:blog_tags;" json:"tags"`
}
