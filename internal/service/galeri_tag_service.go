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

type GaleriTagService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.GaleriTagResponse, error)
	Create(ctx context.Context, req *dto.GaleriTagCreate) (dto.GaleriTagResponse, error)
	Update(ctx context.Context, req *dto.GaleriTagUpdate) error
	Delete(ctx context.Context, id int) error
}

type GaleriTagServiceImpl struct {
	GaleriTagRepository repository.GaleriTagRepository
	DB                  *sql.DB
}

func NewGaleriTagService(donationCategoryRepository repository.GaleriTagRepository, db *sql.DB) GaleriTagService {
	return &GaleriTagServiceImpl{
		GaleriTagRepository: donationCategoryRepository,
		DB:                  db,
	}
}

func (s *GaleriTagServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.GaleriTagRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.GaleriTagRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToGaleriTagResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *GaleriTagServiceImpl) FindById(ctx context.Context, Id int) (dto.GaleriTagResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.GaleriTagRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.GaleriTagResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.GaleriTagResponse{}, err
		}

		return helper.ToGaleriTagResponse(blocat), nil
	})
	return res.(dto.GaleriTagResponse), err
}

func (s *GaleriTagServiceImpl) Create(ctx context.Context, req *dto.GaleriTagCreate) (dto.GaleriTagResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.GaleriTagResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.GaleriTag{
		Title: req.Title,
		Slug:  req.Slug,
	}

	// save datacategory
	data, err = s.GaleriTagRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.GaleriTagResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.GaleriTagResponse{}, errors.New("error commit transaction")
	}

	return helper.ToGaleriTagResponse(data), nil
}

func (s *GaleriTagServiceImpl) Update(ctx context.Context, req *dto.GaleriTagUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.GaleriTagRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.GaleriTag{
			Id:    req.Id,
			Title: req.Title,
			Slug:  req.Slug,
		}

		// update data category donation
		if err = s.GaleriTagRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *GaleriTagServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.GaleriTagRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.GaleriTagRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
