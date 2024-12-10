package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type BlogService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.BlogResponse, error)
	FindBySlug(ctx context.Context, slug string) (dto.BlogResponse, error)
	Create(ctx context.Context, req *dto.BlogCreate) (dto.BlogResponse, error)
	Update(ctx context.Context, req *dto.BlogUpdate) error
	Delete(ctx context.Context, id int) error
}

type BlogServiceImpl struct {
	BlogRepository repository.BlogRepository
	DB             *sql.DB
}

func NewBlogService(blogRepository repository.BlogRepository, db *sql.DB) BlogService {
	return &BlogServiceImpl{
		BlogRepository: blogRepository,
		DB:             db,
	}
}

func (s *BlogServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blogs, err := s.BlogRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.BlogRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToBlogResponses(blogs),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *BlogServiceImpl) FindById(ctx context.Context, Id int) (dto.BlogResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blog, err := s.BlogRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.BlogResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.BlogResponse{}, err
		}

		return helper.ToBlogResponse(blog), nil
	})
	return res.(dto.BlogResponse), err
}

func (s *BlogServiceImpl) FindBySlug(ctx context.Context, slug string) (dto.BlogResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blog, err := s.BlogRepository.FindBySlug(ctx, tx, slug)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.BlogResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.BlogResponse{}, err
		}

		return helper.ToBlogResponse(blog), nil
	})
	return res.(dto.BlogResponse), err
}

func (s *BlogServiceImpl) Create(ctx context.Context, req *dto.BlogCreate) (dto.BlogResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.BlogResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	blog := entity.Blog{
		Title:          req.Title,
		Slug:           req.Slug,
		Image:          req.Image,
		Content:        req.Content,
		UserId:         req.UserId,
		Author:         req.Author,
		BlogCategoryId: req.BlogCategoryId,
	}

	// save blog
	blog, err = s.BlogRepository.Save(ctx, tx, &blog)
	if err != nil {
		return dto.BlogResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.BlogResponse{}, errors.New("error commit transaction")
	}

	return helper.ToBlogResponse(blog), nil
}

func (s *BlogServiceImpl) Update(ctx context.Context, req *dto.BlogUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check blog by id
		_, err := s.BlogRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		blog_update := entity.Blog{
			Id:             req.Id,
			Slug:           req.Slug,
			Image:          req.Image,
			Content:        req.Content,
			UserId:         req.UserId,
			Author:         req.Author,
			BlogCategoryId: req.BlogCategoryId,
		}

		// update blog data
		if err = s.BlogRepository.Update(ctx, tx, &blog_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *BlogServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check blog by id
		blog, err := s.BlogRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete blog
		if err = s.BlogRepository.Delete(ctx, tx, &blog); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
