package dto

type ListReq struct {
	Page     int `json:"page" form:"page" binding:"required,min=1"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"required,min=1,max=100"`

	SortBy string `json:"sortBy" form:"sortBy" binding:"omitempty,oneof=createdAt updatedAt"`
	Order  string `json:"order" form:"order" binding:"omitempty,oneof=asc desc"`
}
