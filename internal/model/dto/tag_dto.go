package dto

import "github.com/aifuxi/fgo/internal/model"

type TagCreateReq struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

type TagListReq struct {
	ListReq
	Name string `json:"name" form:"name" binding:"omitempty"`
	Slug string `json:"slug" form:"slug" binding:"omitempty"`
}

type TagListResp struct {
	Total int64        `json:"total"`
	Lists []*model.Tag `json:"lists"`
}

type TagFindByIDReq struct {
	ID int64 `uri:"id" binding:"required"`
}

type TagUpdateReq struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}
