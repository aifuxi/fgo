package repository

import (
	"context"
	"errors"

	"github.com/aifuxi/fgo/internal/model"
	"gorm.io/gorm"
)

type UserListOption struct {
	Page     int
	PageSize int
	Nickname string
	Email    string
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	List(ctx context.Context, option UserListOption) ([]*model.User, int64, error)
	FindByID(ctx context.Context, id uint) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	DeleteByID(ctx context.Context, id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Model(user).Association("Roles").Replace(user.Roles)
	// Also update user fields
	// return r.db.WithContext(ctx).Save(user).Error
	// Better to use Updates for fields and Association for roles, but Save works if ID is present.
	// Let's stick to Updates for fields and handle Roles separately if needed, but Save is easier for full update.
	// However, usually we update specific fields.
	// Let's assume the service handles the logic of what to update.
	// But here, let's just save the user.
	return r.db.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Save(user).Error
}

func (r *userRepo) List(ctx context.Context, option UserListOption) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	query := r.db.WithContext(ctx).Model(&model.User{})

	if option.Nickname != "" {
		query = query.Where("nickname LIKE ?", "%"+option.Nickname+"%")
	}

	if option.Email != "" {
		query = query.Where("email LIKE ?", "%"+option.Email+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("Roles").Preload("Roles.Permissions").
		Offset((option.Page - 1) * option.PageSize).
		Limit(option.PageSize).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepo) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Roles").Preload("Roles.Permissions").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Roles").Preload("Roles.Permissions").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) DeleteByID(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}
