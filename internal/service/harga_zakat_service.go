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

type HargaZakatService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.HargaZakatResponse, error)
	Create(ctx context.Context, req *dto.HargaZakatCreate) (dto.HargaZakatResponse, error)
	Update(ctx context.Context, req *dto.HargaZakatUpdate) error
	Delete(ctx context.Context, id int) error
}

type HargaZakatServiceImpl struct {
	HargaZakatRepository repository.HargaZakatRepository
	DB                   *sql.DB
}

func NewHargaZakatService(donationCategoryRepository repository.HargaZakatRepository, db *sql.DB) HargaZakatService {
	return &HargaZakatServiceImpl{
		HargaZakatRepository: donationCategoryRepository,
		DB:                   db,
	}
}

func (s *HargaZakatServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.HargaZakatRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.HargaZakatRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToHargaZakatResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *HargaZakatServiceImpl) FindById(ctx context.Context, Id int) (dto.HargaZakatResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.HargaZakatRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.HargaZakatResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.HargaZakatResponse{}, err
		}

		return helper.ToHargaZakatResponse(blocat), nil
	})
	return res.(dto.HargaZakatResponse), err
}

func (s *HargaZakatServiceImpl) Create(ctx context.Context, req *dto.HargaZakatCreate) (dto.HargaZakatResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.HargaZakatResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.HargaZakat{
		DonationListId: req.DonationListId,
		Title:          req.Title,
		Price:          req.Price,
	}

	// save datacategory
	data, err = s.HargaZakatRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.HargaZakatResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.HargaZakatResponse{}, errors.New("error commit transaction")
	}

	return helper.ToHargaZakatResponse(data), nil
}

func (s *HargaZakatServiceImpl) Update(ctx context.Context, req *dto.HargaZakatUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.HargaZakatRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.HargaZakat{
			Id:             req.Id,
			DonationListId: req.DonationListId,
			Title:          req.Title,
			Price:          req.Price,
		}

		// update data category donation
		if err = s.HargaZakatRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *HargaZakatServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.HargaZakatRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.HargaZakatRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
