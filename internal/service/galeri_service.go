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

type GaleriService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.GaleriResponse, error)
	Create(ctx context.Context, req *dto.GaleriCreate) (dto.GaleriResponse, error)
	Update(ctx context.Context, req *dto.GaleriUpdate) error
	Delete(ctx context.Context, id int) error
}

type GaleriServiceImpl struct {
	GaleriRepository repository.GaleriRepository
	DB               *sql.DB
}

func NewGaleriService(donationCategoryRepository repository.GaleriRepository, db *sql.DB) GaleriService {
	return &GaleriServiceImpl{
		GaleriRepository: donationCategoryRepository,
		DB:               db,
	}
}

func (s *GaleriServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.GaleriRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.GaleriRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToGaleriResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *GaleriServiceImpl) FindById(ctx context.Context, Id int) (dto.GaleriResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.GaleriRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.GaleriResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.GaleriResponse{}, err
		}

		return helper.ToGaleriResponse(blocat), nil
	})
	return res.(dto.GaleriResponse), err
}

func (s *GaleriServiceImpl) Create(ctx context.Context, req *dto.GaleriCreate) (dto.GaleriResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.GaleriResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.Galeri{
		Title:       req.Title,
		Slug:        req.Slug,
		Image:       req.Image,
		GaleryTagId: req.GaleryTagId,
	}

	// save datacategory
	data, err = s.GaleriRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.GaleriResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.GaleriResponse{}, errors.New("error commit transaction")
	}

	return helper.ToGaleriResponse(data), nil
}

func (s *GaleriServiceImpl) Update(ctx context.Context, req *dto.GaleriUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.GaleriRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.Galeri{
			Id:          req.Id,
			Title:       req.Title,
			Slug:        req.Slug,
			Image:       req.Image,
			GaleryTagId: req.GaleryTagId,
		}

		// update data category donation
		if err = s.GaleriRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *GaleriServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.GaleriRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.GaleriRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
