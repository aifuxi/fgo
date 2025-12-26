package dto

import "github.com/aifuxi/fgo/internal/model"

type RoleCreateReq struct {
	Name        string `json:"name" binding:"required,max=50"`
	Code        string `json:"code" binding:"required,max=50"`
	Description string `json:"description" binding:"max=255"`
}

type RoleUpdateReq struct {
	Name        string `json:"name" binding:"omitempty,max=50"`
	Code        string `json:"code" binding:"omitempty,max=50"`
	Description string `json:"description" binding:"max=255"`
}

type RoleListReq struct {
	Page     int    `json:"page" form:"page" binding:"min=1"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"min=1,max=100"`
	Name     string `json:"name" form:"name"`
	Code     string `json:"code" form:"code"`
}

type RoleResp struct {
	model.CommonModel
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type RoleListResp struct {
	Total int64       `json:"total"`
	Lists []*RoleResp `json:"lists"`
}
