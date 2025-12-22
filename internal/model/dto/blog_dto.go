package dto

import "github.com/aifuxi/fgo/internal/model"

type BlogCreateReq struct {
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description" binding:"required"`
	Cover       string `json:"cover"`
	Content     string `json:"content" binding:"required"`
	Published   bool   `json:"published"`
	CategoryID  uint   `json:"categoryID"`
	TagIDs      []uint `json:"tagIDs"`
}

type BlogListReq struct {
	ListReq
	Title           string `json:"title" form:"title" binding:"omitempty"`
	Slug            string `json:"slug" form:"slug" binding:"omitempty"`
	PublishedStatus string `json:"publishedStatus" form:"published" binding:"omitempty,oneof=published unpublished"`
	CategoryID      uint   `json:"categoryID" form:"categoryID" binding:"omitempty"`
	TagIDs          []uint `json:"tagIDs" form:"tagIDs" binding:"omitempty"`
}

type BlogListResp struct {
	Total int64         `json:"total"`
	Lists []*model.Blog `json:"lists"`
}

type BlogFindByIDReq struct {
	ID uint `uri:"id" binding:"required"`
}

type BlogUpdateReq struct {
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description" binding:"required"`
	Cover       string `json:"cover"`
	Content     string `json:"content" binding:"required"`
	Published   bool   `json:"published"`
	CategoryID  uint   `json:"categoryID"`
	TagIDs      []uint `json:"tagIDs"`
}
