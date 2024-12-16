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

type DonationService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.DonationResponse, error)
	Create(ctx context.Context, req *dto.DonationCreate) (dto.DonationResponse, error)
	Update(ctx context.Context, req *dto.DonationUpdate) error
	Delete(ctx context.Context, id int) error
}

type DonationServiceImpl struct {
	DonationRepository repository.DonationRepository
	DB                 *sql.DB
}

func NewDonationService(donationCategoryRepository repository.DonationRepository, db *sql.DB) DonationService {
	return &DonationServiceImpl{
		DonationRepository: donationCategoryRepository,
		DB:                 db,
	}
}

func (s *DonationServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.DonationRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.DonationRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToDonationResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *DonationServiceImpl) FindById(ctx context.Context, Id int) (dto.DonationResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.DonationRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.DonationResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.DonationResponse{}, err
		}

		return helper.ToDonationResponse(blocat), nil
	})
	return res.(dto.DonationResponse), err
}

func (s *DonationServiceImpl) Create(ctx context.Context, req *dto.DonationCreate) (dto.DonationResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.DonationResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.Donation{
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		Amount:         req.Amount,
		Status:         req.Status,
		Reference:      req.Reference,
		SnapToken:      req.SnapToken,
		DonationListId: req.DonationListId,
		CharityListId:  req.CharityListId,
		UserId:         req.UserId,
		OrderId:        req.OrderId,
	}

	// save datacategory
	data, err = s.DonationRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.DonationResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.DonationResponse{}, errors.New("error commit transaction")
	}

	return helper.ToDonationResponse(data), nil
}

func (s *DonationServiceImpl) Update(ctx context.Context, req *dto.DonationUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.DonationRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.Donation{
			Id:             req.Id,
			Name:           req.Name,
			Email:          req.Email,
			Phone:          req.Phone,
			Amount:         req.Amount,
			Status:         req.Status,
			Reference:      req.Reference,
			SnapToken:      req.SnapToken,
			DonationListId: req.DonationListId,
			CharityListId:  req.CharityListId,
			UserId:         req.UserId,
			OrderId:        req.OrderId,
		}

		// update data category donation
		if err = s.DonationRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *DonationServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.DonationRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.DonationRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
