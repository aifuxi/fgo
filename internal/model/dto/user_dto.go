package dto

import "github.com/aifuxi/fgo/internal/model"

type UserRegisterReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateReq struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"omitempty,email"`
	RoleIDs  []uint `json:"roleIDs"`
}

type UserListReq struct {
	ListReq
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
}

type UserResp struct {
	model.CommonModel
	Nickname string        `json:"nickname"`
	Email    string        `json:"email"`
	Roles    []*model.Role `json:"roles,omitempty"`
}

type UserListResp struct {
	Total int64       `json:"total"`
	Lists []*UserResp `json:"lists"`
}

type UserFindByIDReq struct {
	ID uint `uri:"id" binding:"required"`
}
