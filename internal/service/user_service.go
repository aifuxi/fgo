package service

import (
	"context"
	"errors"

	"github.com/aifuxi/fgo/internal/model"
	"github.com/aifuxi/fgo/internal/model/dto"
	"github.com/aifuxi/fgo/internal/repository"
	"github.com/aifuxi/fgo/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *dto.UserRegisterReq) error
	Login(ctx context.Context, req *dto.UserLoginReq) (string, error)
	Update(ctx context.Context, id uint, req *dto.UserUpdateReq) error
	List(ctx context.Context, req *dto.UserListReq) (*dto.UserListResp, error)
	FindByID(ctx context.Context, id uint) (*dto.UserResp, error)
	DeleteByID(ctx context.Context, id uint) error
}

type userService struct {
	repo     repository.UserRepository
	roleRepo repository.RoleRepository
}

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrUserEmailExists        = errors.New("user email already exists")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrVisitorRoleNotFound    = errors.New("visitor role not found")
)

func NewUserService(repo repository.UserRepository, roleRepo repository.RoleRepository) UserService {
	return &userService{repo: repo, roleRepo: roleRepo}
}

func (s *userService) Register(ctx context.Context, req *dto.UserRegisterReq) error {
	// Check if email exists
	existingUser, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return ErrUserEmailExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Find visitor role
	visitorRole, err := s.roleRepo.FindByCode(ctx, model.RoleCodeVisitor)
	if err != nil {
		return err
	}
	if visitorRole == nil {
		return ErrVisitorRoleNotFound
	}

	user := &model.User{
		Nickname: req.Nickname,
		Email:    req.Email,
		Password: string(hashedPassword),
		Roles: []*model.Role{
			visitorRole,
		},
	}

	return s.repo.Create(ctx, user)
}

func (s *userService) Login(ctx context.Context, req *dto.UserLoginReq) (string, error) {
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrInvalidEmailOrPassword
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", ErrInvalidEmailOrPassword
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Update(ctx context.Context, id uint, req *dto.UserUpdateReq) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	if req.Email != "" && req.Email != user.Email {
		existingUser, err := s.repo.FindByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if existingUser != nil {
			return ErrUserEmailExists
		}
		user.Email = req.Email
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}

	if len(req.RoleIDs) > 0 {
		var roles []*model.Role
		for _, roleID := range req.RoleIDs {
			roles = append(roles, &model.Role{
				CommonModel: model.CommonModel{ID: roleID},
			})
		}
		user.Roles = roles
	}

	return s.repo.Update(ctx, user)
}

func (s *userService) List(ctx context.Context, req *dto.UserListReq) (*dto.UserListResp, error) {
	users, total, err := s.repo.List(ctx, repository.UserListOption{
		Page:     req.Page,
		PageSize: req.PageSize,
		Nickname: req.Nickname,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	userRespList := convertToUserRespList(users)

	return &dto.UserListResp{
		Total: total,
		Lists: userRespList,
	}, nil
}

func convertToUserRespList(users []*model.User) []*dto.UserResp {
	var userRespList []*dto.UserResp
	for _, user := range users {
		userRespList = append(userRespList, &dto.UserResp{
			CommonModel: user.CommonModel,
			Nickname:    user.Nickname,
			Email:       user.Email,
			Roles:       user.Roles,
		})
	}
	return userRespList
}

func (s *userService) FindByID(ctx context.Context, id uint) (*dto.UserResp, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return &dto.UserResp{
		CommonModel: user.CommonModel,
		Nickname:    user.Nickname,
		Email:       user.Email,
		Roles:       user.Roles,
	}, nil
}

func (s *userService) DeleteByID(ctx context.Context, id uint) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}
	return s.repo.DeleteByID(ctx, id)
}
