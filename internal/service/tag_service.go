package service

import (
	"context"
	"errors"

	"github.com/aifuxi/fgo/internal/model"
	"github.com/aifuxi/fgo/internal/model/dto"
	"github.com/aifuxi/fgo/internal/repository"
)

type TagService interface {
	Create(ctx context.Context, req dto.TagCreateReq) (*model.Tag, error)
	List(ctx context.Context, req dto.TagListReq) ([]*model.Tag, int64, error)
	FindByID(ctx context.Context, id int64) (*model.Tag, error)
	UpdateByID(ctx context.Context, id int64, req dto.TagUpdateReq) (*model.Tag, error)
	DeleteByID(ctx context.Context, id int64) error
}

type tagService struct {
	tagRepo repository.TagRepository
}

var (
	ErrTagNotFound   = errors.New("tag not found")
	ErrTagSlugExists = errors.New("tag slug already exists")
	ErrTagNameExists = errors.New("tag name already exists")
)

func NewTagService(tagRepo repository.TagRepository) TagService {
	return &tagService{tagRepo: tagRepo}
}

func (s *tagService) Create(ctx context.Context, req dto.TagCreateReq) (*model.Tag, error) {
	tag := &model.Tag{
		Name: req.Name,
		Slug: req.Slug,
	}

	existNameTag, err := s.tagRepo.FindByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	if existNameTag.ID != 0 {
		return nil, ErrTagNameExists
	}

	existSlugTag, err := s.tagRepo.FindBySlug(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	if existSlugTag.ID != 0 {
		return nil, ErrTagSlugExists
	}

	createdTag, err := s.tagRepo.Create(ctx, tag)
	if err != nil {
		return nil, err
	}

	return createdTag, nil
}

func (s *tagService) List(ctx context.Context, req dto.TagListReq) ([]*model.Tag, int64, error) {
	var tags []*model.Tag
	var total int64
	var err error

	tags, total, err = s.tagRepo.List(ctx, repository.TagListOption{
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

func (s *tagService) FindByID(ctx context.Context, id int64) (*model.Tag, error) {
	tag, err := s.tagRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if tag.ID == 0 {
		return nil, ErrTagNotFound
	}

	return tag, nil
}

func (s *tagService) DeleteByID(ctx context.Context, id int64) error {
	tag, err := s.tagRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if tag.ID == 0 {
		return ErrTagNotFound
	}

	return s.tagRepo.DeleteByID(ctx, id)
}

func (s *tagService) UpdateByID(ctx context.Context, id int64, req dto.TagUpdateReq) (*model.Tag, error) {
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
