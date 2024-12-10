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

type BlogCategoryService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.BlogCategoryResponse, error)
	Create(ctx context.Context, req *dto.BlogCategoryCreate) (dto.BlogCategoryResponse, error)
	Update(ctx context.Context, req *dto.BlogCategoryUpdate) error
	Delete(ctx context.Context, id int) error
}

type BlogCategoryServiceImpl struct {
	BlogCategoryRepository repository.BlogCategoryRepository
	DB                     *sql.DB
}

func NewBlogCategoryService(blogCategoryRepository repository.BlogCategoryRepository, db *sql.DB) BlogCategoryService {
	return &BlogCategoryServiceImpl{
		BlogCategoryRepository: blogCategoryRepository,
		DB:                     db,
	}
}

func (s *BlogCategoryServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blogs, err := s.BlogCategoryRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.BlogCategoryRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToBlogCategoryResponses(blogs),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *BlogCategoryServiceImpl) FindById(ctx context.Context, Id int) (dto.BlogCategoryResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.BlogCategoryRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.BlogCategoryResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.BlogCategoryResponse{}, err
		}

		return helper.ToBlogCategoryResponse(blocat), nil
	})
	return res.(dto.BlogCategoryResponse), err
}

func (s *BlogCategoryServiceImpl) Create(ctx context.Context, req *dto.BlogCategoryCreate) (dto.BlogCategoryResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.BlogCategoryResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	blog := entity.BlogCategory{
		Title:       req.Title,
		Description: req.Description,
	}

	// save blogcategory
	blog, err = s.BlogCategoryRepository.Save(ctx, tx, &blog)
	if err != nil {
		return dto.BlogCategoryResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.BlogCategoryResponse{}, errors.New("error commit transaction")
	}

	return helper.ToBlogCategoryResponse(blog), nil
}

func (s *BlogCategoryServiceImpl) Update(ctx context.Context, req *dto.BlogCategoryUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check blog category by id
		_, err := s.BlogCategoryRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		blog_update := entity.BlogCategory{
			Id:          req.Id,
			Title:       req.Title,
			Description: req.Description,
		}

		// update blog category data
		if err = s.BlogCategoryRepository.Update(ctx, tx, &blog_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *BlogCategoryServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check blog category by id
		blogcat, err := s.BlogCategoryRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete blog category
		if err = s.BlogCategoryRepository.Delete(ctx, tx, &blogcat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
