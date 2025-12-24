package handler

import (
	"errors"
	"strconv"

	"github.com/aifuxi/fgo/internal/model/dto"
	"github.com/aifuxi/fgo/internal/service"
	"github.com/aifuxi/fgo/pkg/response"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	svc service.RoleService
}

func NewRoleHandler(svc service.RoleService) *RoleHandler {
	return &RoleHandler{svc: svc}
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req dto.RoleCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	if err := h.svc.Create(c, &req); err != nil {
		if errors.Is(err, service.ErrRoleNameExists) || errors.Is(err, service.ErrRoleCodeExists) {
			response.BusinessError(c, err.Error())
			return
		}
		response.BusinessError(c, "Failed to create role")
		return
	}

	response.Success(c, nil)
}

func (h *RoleHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ParamError(c, "Invalid role ID")
		return
	}

	var req dto.RoleUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	if err := h.svc.Update(c, int64(id), &req); err != nil {
		if errors.Is(err, service.ErrRoleNotFound) {
			response.BusinessError(c, err.Error())
			return
		}
		if errors.Is(err, service.ErrRoleNameExists) || errors.Is(err, service.ErrRoleCodeExists) {
			response.BusinessError(c, err.Error())
			return
		}
		response.BusinessError(c, "Failed to update role")
		return
	}

	response.Success(c, nil)
}

func (h *RoleHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ParamError(c, "Invalid role ID")
		return
	}

	if err := h.svc.Delete(c, int64(id)); err != nil {
		if errors.Is(err, service.ErrRoleNotFound) {
			response.BusinessError(c, err.Error())
			return
		}
		response.BusinessError(c, "Failed to delete role")
		return
	}

	response.Success(c, nil)
}

func (h *RoleHandler) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ParamError(c, "Invalid role ID")
		return
	}

	role, err := h.svc.FindByID(c, int64(id))
	if err != nil {
		if errors.Is(err, service.ErrRoleNotFound) {
			response.BusinessError(c, err.Error())
			return
		}
		response.BusinessError(c, "Failed to get role")
		return
	}

	response.Success(c, role)
}

func (h *RoleHandler) List(c *gin.Context) {
	var req dto.RoleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	resp, err := h.svc.List(c, &req)
	if err != nil {
		response.BusinessError(c, "Failed to list roles")
		return
	}

	response.Success(c, dto.RoleListResp{
		Total: resp.Total,
		Lists: resp.Lists,
	})
}
