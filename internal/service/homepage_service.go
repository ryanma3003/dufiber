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

type HomepageService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.HomepageResponse, error)
	Create(ctx context.Context, req *dto.HomepageCreate) (dto.HomepageResponse, error)
	Update(ctx context.Context, req *dto.HomepageUpdate) error
	Delete(ctx context.Context, id int) error
}

type HomepageServiceImpl struct {
	HomepageRepository repository.HomepageRepository
	DB                 *sql.DB
}

func NewHomepageService(donationCategoryRepository repository.HomepageRepository, db *sql.DB) HomepageService {
	return &HomepageServiceImpl{
		HomepageRepository: donationCategoryRepository,
		DB:                 db,
	}
}

func (s *HomepageServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.HomepageRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.HomepageRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToHomepageResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *HomepageServiceImpl) FindById(ctx context.Context, Id int) (dto.HomepageResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.HomepageRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.HomepageResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.HomepageResponse{}, err
		}

		return helper.ToHomepageResponse(blocat), nil
	})
	return res.(dto.HomepageResponse), err
}

func (s *HomepageServiceImpl) Create(ctx context.Context, req *dto.HomepageCreate) (dto.HomepageResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.HomepageResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.Homepage{
		MainImage:       req.MainImage,
		MainText:        req.MainText,
		MainTitle:       req.MainTitle,
		KalkulatorTitle: req.KalkulatorTitle,
		KalkulatorText:  req.KalkulatorText,
		PersText:        req.PersText,
		PublikasiText:   req.PublikasiText,
	}

	// save datacategory
	data, err = s.HomepageRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.HomepageResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.HomepageResponse{}, errors.New("error commit transaction")
	}

	return helper.ToHomepageResponse(data), nil
}

func (s *HomepageServiceImpl) Update(ctx context.Context, req *dto.HomepageUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.HomepageRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.Homepage{
			Id:              req.Id,
			MainImage:       req.MainImage,
			MainText:        req.MainText,
			MainTitle:       req.MainTitle,
			KalkulatorTitle: req.KalkulatorTitle,
			KalkulatorText:  req.KalkulatorText,
			PersText:        req.PersText,
			PublikasiText:   req.PublikasiText,
		}

		// update data category donation
		if err = s.HomepageRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *HomepageServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.HomepageRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.HomepageRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
