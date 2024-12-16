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

type PrivacyService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.PrivacyResponse, error)
	Create(ctx context.Context, req *dto.PrivacyCreate) (dto.PrivacyResponse, error)
	Update(ctx context.Context, req *dto.PrivacyUpdate) error
	Delete(ctx context.Context, id int) error
}

type PrivacyServiceImpl struct {
	PrivacyRepository repository.PrivacyRepository
	DB                *sql.DB
}

func NewPrivacyService(donationCategoryRepository repository.PrivacyRepository, db *sql.DB) PrivacyService {
	return &PrivacyServiceImpl{
		PrivacyRepository: donationCategoryRepository,
		DB:                db,
	}
}

func (s *PrivacyServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.PrivacyRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.PrivacyRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToPrivacyResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *PrivacyServiceImpl) FindById(ctx context.Context, Id int) (dto.PrivacyResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.PrivacyRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.PrivacyResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.PrivacyResponse{}, err
		}

		return helper.ToPrivacyResponse(blocat), nil
	})
	return res.(dto.PrivacyResponse), err
}

func (s *PrivacyServiceImpl) Create(ctx context.Context, req *dto.PrivacyCreate) (dto.PrivacyResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.PrivacyResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.Privacy{
		Title: req.Title,
		Text:  req.Text,
	}

	// save datacategory
	data, err = s.PrivacyRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.PrivacyResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.PrivacyResponse{}, errors.New("error commit transaction")
	}

	return helper.ToPrivacyResponse(data), nil
}

func (s *PrivacyServiceImpl) Update(ctx context.Context, req *dto.PrivacyUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.PrivacyRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.Privacy{
			Id:    req.Id,
			Title: req.Title,
			Text:  req.Text,
		}

		// update data category donation
		if err = s.PrivacyRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *PrivacyServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.PrivacyRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.PrivacyRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
