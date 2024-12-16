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

type DonationListService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.DonationListResponse, error)
	Create(ctx context.Context, req *dto.DonationListCreate) (dto.DonationListResponse, error)
	Update(ctx context.Context, req *dto.DonationListUpdate) error
	Delete(ctx context.Context, id int) error
}

type DonationListServiceImpl struct {
	DonationListRepository repository.DonationListRepository
	DB                     *sql.DB
}

func NewDonationListService(donationCategoryRepository repository.DonationListRepository, db *sql.DB) DonationListService {
	return &DonationListServiceImpl{
		DonationListRepository: donationCategoryRepository,
		DB:                     db,
	}
}

func (s *DonationListServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.DonationListRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.DonationListRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToDonationListResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *DonationListServiceImpl) FindById(ctx context.Context, Id int) (dto.DonationListResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.DonationListRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.DonationListResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.DonationListResponse{}, err
		}

		return helper.ToDonationListResponse(blocat), nil
	})
	return res.(dto.DonationListResponse), err
}

func (s *DonationListServiceImpl) Create(ctx context.Context, req *dto.DonationListCreate) (dto.DonationListResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.DonationListResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.DonationList{
		Title:              req.Title,
		Description:        req.Description,
		Code:               req.Code,
		DonationCategoryId: req.DonationCategoryId,
	}

	// save datacategory
	data, err = s.DonationListRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.DonationListResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.DonationListResponse{}, errors.New("error commit transaction")
	}

	return helper.ToDonationListResponse(data), nil
}

func (s *DonationListServiceImpl) Update(ctx context.Context, req *dto.DonationListUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.DonationListRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.DonationList{
			Id:                 req.Id,
			Title:              req.Title,
			Description:        req.Description,
			Code:               req.Code,
			DonationCategoryId: req.DonationCategoryId,
		}

		// update data category donation
		if err = s.DonationListRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *DonationListServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.DonationListRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.DonationListRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
