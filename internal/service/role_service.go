package service

import (
	"context"
	"errors"

	"github.com/aifuxi/fgo/internal/model"
	"github.com/aifuxi/fgo/internal/model/dto"
	"github.com/aifuxi/fgo/internal/repository"
)

var (
	ErrRoleNotFound   = errors.New("role not found")
	ErrRoleNameExists = errors.New("role name already exists")
	ErrRoleCodeExists = errors.New("role code already exists")
)

type RoleService interface {
	Create(ctx context.Context, req *dto.RoleCreateReq) error
	Update(ctx context.Context, id uint, req *dto.RoleUpdateReq) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*dto.RoleResp, error)
	List(ctx context.Context, req *dto.RoleListReq) (*dto.RoleListResp, error)
}

type roleService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) Create(ctx context.Context, req *dto.RoleCreateReq) error {
	// Check if name exists
	existingRole, err := s.repo.FindByName(ctx, req.Name)
	if err != nil {
		return err
	}
	if existingRole != nil {
		return ErrRoleNameExists
	}

	// Check if code exists
	existingRole, err = s.repo.FindByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if existingRole != nil {
		return ErrRoleCodeExists
	}

	role := &model.Role{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}

	return s.repo.Create(ctx, role)
}

func (s *roleService) Update(ctx context.Context, id uint, req *dto.RoleUpdateReq) error {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if role == nil {
		return ErrRoleNotFound
	}

	if req.Name != "" && req.Name != role.Name {
		existingRole, err := s.repo.FindByName(ctx, req.Name)
		if err != nil {
			return err
		}
		if existingRole != nil {
			return ErrRoleNameExists
		}
		role.Name = req.Name
	}

	if req.Code != "" && req.Code != role.Code {
		existingRole, err := s.repo.FindByCode(ctx, req.Code)
		if err != nil {
			return err
		}
		if existingRole != nil {
			return ErrRoleCodeExists
		}
		role.Code = req.Code
	}

	if req.Description != "" {
		role.Description = req.Description
	}

	return s.repo.Update(ctx, role)
}

func (s *roleService) Delete(ctx context.Context, id uint) error {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if role == nil {
		return ErrRoleNotFound
	}

	return s.repo.Delete(ctx, id)
}

func (s *roleService) FindByID(ctx context.Context, id uint) (*dto.RoleResp, error) {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, ErrRoleNotFound
	}

	return &dto.RoleResp{
		CommonModel: role.CommonModel,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
	}, nil
}

func (s *roleService) List(ctx context.Context, req *dto.RoleListReq) (*dto.RoleListResp, error) {
	roles, total, err := s.repo.List(ctx, repository.RoleListOption{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}

	var list []*dto.RoleResp
	for _, role := range roles {
		list = append(list, &dto.RoleResp{
			CommonModel: role.CommonModel,
			Name:        role.Name,
			Code:        role.Code,
			Description: role.Description,
		})
	}

	return &dto.RoleListResp{
		Total: total,
		Lists: list,
	}, nil
}
