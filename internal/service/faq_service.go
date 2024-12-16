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

type FaqService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.FaqResponse, error)
	Create(ctx context.Context, req *dto.FaqCreate) (dto.FaqResponse, error)
	Update(ctx context.Context, req *dto.FaqUpdate) error
	Delete(ctx context.Context, id int) error
}

type FaqServiceImpl struct {
	FaqRepository repository.FaqRepository
	DB            *sql.DB
}

func NewFaqService(donationCategoryRepository repository.FaqRepository, db *sql.DB) FaqService {
	return &FaqServiceImpl{
		FaqRepository: donationCategoryRepository,
		DB:            db,
	}
}

func (s *FaqServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.FaqRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.FaqRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToFaqResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *FaqServiceImpl) FindById(ctx context.Context, Id int) (dto.FaqResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.FaqRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.FaqResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.FaqResponse{}, err
		}

		return helper.ToFaqResponse(blocat), nil
	})
	return res.(dto.FaqResponse), err
}

func (s *FaqServiceImpl) Create(ctx context.Context, req *dto.FaqCreate) (dto.FaqResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.FaqResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.Faq{
		Pertanyaan: req.Pertanyaan,
		Jawaban:    req.Jawaban,
	}

	// save datacategory
	data, err = s.FaqRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.FaqResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.FaqResponse{}, errors.New("error commit transaction")
	}

	return helper.ToFaqResponse(data), nil
}

func (s *FaqServiceImpl) Update(ctx context.Context, req *dto.FaqUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.FaqRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.Faq{
			Id:         req.Id,
			Pertanyaan: req.Pertanyaan,
			Jawaban:    req.Jawaban,
		}

		// update data category donation
		if err = s.FaqRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *FaqServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.FaqRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.FaqRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
