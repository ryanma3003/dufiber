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

type TermService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.TermResponse, error)
	Create(ctx context.Context, req *dto.TermCreate) (dto.TermResponse, error)
	Update(ctx context.Context, req *dto.TermUpdate) error
	Delete(ctx context.Context, id int) error
}

type TermServiceImpl struct {
	TermRepository repository.TermRepository
	DB             *sql.DB
}

func NewTermService(donationCategoryRepository repository.TermRepository, db *sql.DB) TermService {
	return &TermServiceImpl{
		TermRepository: donationCategoryRepository,
		DB:             db,
	}
}

func (s *TermServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.TermRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.TermRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToTermResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *TermServiceImpl) FindById(ctx context.Context, Id int) (dto.TermResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.TermRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.TermResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.TermResponse{}, err
		}

		return helper.ToTermResponse(blocat), nil
	})
	return res.(dto.TermResponse), err
}

func (s *TermServiceImpl) Create(ctx context.Context, req *dto.TermCreate) (dto.TermResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.TermResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.Term{
		Title: req.Title,
		Text:  req.Text,
	}

	// save datacategory
	data, err = s.TermRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.TermResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.TermResponse{}, errors.New("error commit transaction")
	}

	return helper.ToTermResponse(data), nil
}

func (s *TermServiceImpl) Update(ctx context.Context, req *dto.TermUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.TermRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.Term{
			Id:    req.Id,
			Title: req.Title,
			Text:  req.Text,
		}

		// update data category donation
		if err = s.TermRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *TermServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.TermRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.TermRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
