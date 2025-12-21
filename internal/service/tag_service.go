package service

import (
	"context"
	"errors"

	"github.com/aifuxi/fgo/internal/model"
	"github.com/aifuxi/fgo/internal/model/dto"
	"github.com/aifuxi/fgo/internal/repository"
)

type ITagService interface {
	Create(ctx context.Context, req dto.TagCreateReq) (*model.Tag, error)
	FindAll(ctx context.Context, req dto.TagListReq) ([]*model.Tag, int64, error)
	FindByID(ctx context.Context, id uint) (*model.Tag, error)
	DeleteByID(ctx context.Context, id uint) error
	UpdateByID(ctx context.Context, id uint, req dto.TagUpdateReq) (*model.Tag, error)
}

type TagService struct {
	tagRepo repository.ITagRepository
}

var (
	ErrTagNotFound = errors.New("tag not found")
)

func NewTagService(tagRepo repository.ITagRepository) ITagService {
	return &TagService{tagRepo: tagRepo}
}

func (s *TagService) Create(ctx context.Context, req dto.TagCreateReq) (*model.Tag, error) {
	tag := &model.Tag{
		Name: req.Name,
		Slug: req.Slug,
	}

	existNameTag, err := s.tagRepo.FindByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	if existNameTag.ID != 0 {
		return nil, errors.New("tag name already exists")
	}

	existSlugTag, err := s.tagRepo.FindBySlug(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	if existSlugTag.ID != 0 {
		return nil, errors.New("tag slug already exists")
	}

	createdTag, err := s.tagRepo.Create(ctx, tag)
	if err != nil {
		return nil, err
	}

	return createdTag, nil
}

func (s *TagService) FindAll(ctx context.Context, req dto.TagListReq) ([]*model.Tag, int64, error) {
	var tags []*model.Tag
	var total int64
	var err error

	tags, total, err = s.tagRepo.FindAll(ctx, repository.TagListOption{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		Slug:     req.Slug,
		SortBy:   req.SortBy,
		Order:    req.Order,
	})
	if err != nil {
		return nil, 0, err
	}

	return tags, total, nil
}

func (s *TagService) FindByID(ctx context.Context, id uint) (*model.Tag, error) {
	tag, err := s.tagRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if tag.ID == 0 {
		return nil, ErrTagNotFound
	}

	return tag, nil
}

func (s *TagService) DeleteByID(ctx context.Context, id uint) error {
	tag, err := s.tagRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if tag.ID == 0 {
		return ErrTagNotFound
	}

	return s.tagRepo.DeleteByID(ctx, id)
}

func (s *TagService) UpdateByID(ctx context.Context, id uint, req dto.TagUpdateReq) (*model.Tag, error) {
	tag, err := s.tagRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if tag.ID == 0 {
		return nil, ErrTagNotFound
	}

	tag.Name = req.Name
	tag.Slug = req.Slug

	updatedTag, err := s.tagRepo.UpdateByID(ctx, id, tag)
	if err != nil {
		return nil, err
	}

	return updatedTag, nil
}
