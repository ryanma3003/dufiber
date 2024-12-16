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

type DonationCategoryService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.DonationCategoryResponse, error)
	Create(ctx context.Context, req *dto.DonationCategoryCreate) (dto.DonationCategoryResponse, error)
	Update(ctx context.Context, req *dto.DonationCategoryUpdate) error
	Delete(ctx context.Context, id int) error
}

type DonationCategoryServiceImpl struct {
	DonationCategoryRepository repository.DonationCategoryRepository
	DB                         *sql.DB
}

func NewDonationCategoryService(donationCategoryRepository repository.DonationCategoryRepository, db *sql.DB) DonationCategoryService {
	return &DonationCategoryServiceImpl{
		DonationCategoryRepository: donationCategoryRepository,
		DB:                         db,
	}
}

func (s *DonationCategoryServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.DonationCategoryRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.DonationCategoryRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToDonationCategoryResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *DonationCategoryServiceImpl) FindById(ctx context.Context, Id int) (dto.DonationCategoryResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.DonationCategoryRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.DonationCategoryResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.DonationCategoryResponse{}, err
		}

		return helper.ToDonationCategoryResponse(blocat), nil
	})
	return res.(dto.DonationCategoryResponse), err
}

func (s *DonationCategoryServiceImpl) Create(ctx context.Context, req *dto.DonationCategoryCreate) (dto.DonationCategoryResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.DonationCategoryResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.DonationCategory{
		Title:       req.Title,
		Description: req.Description,
	}

	// save datacategory
	data, err = s.DonationCategoryRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.DonationCategoryResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.DonationCategoryResponse{}, errors.New("error commit transaction")
	}

	return helper.ToDonationCategoryResponse(data), nil
}

func (s *DonationCategoryServiceImpl) Update(ctx context.Context, req *dto.DonationCategoryUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.DonationCategoryRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.DonationCategory{
			Id:          req.Id,
			Title:       req.Title,
			Description: req.Description,
		}

		// update data category donation
		if err = s.DonationCategoryRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *DonationCategoryServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.DonationCategoryRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.DonationCategoryRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
