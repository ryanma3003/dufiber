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

type AboutService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.AboutResponse, error)
	Create(ctx context.Context, req *dto.AboutCreate) (dto.AboutResponse, error)
	Update(ctx context.Context, req *dto.AboutUpdate) error
	Delete(ctx context.Context, id int) error
}

type AboutServiceImpl struct {
	AboutRepository repository.AboutRepository
	DB              *sql.DB
}

func NewAboutService(donationCategoryRepository repository.AboutRepository, db *sql.DB) AboutService {
	return &AboutServiceImpl{
		AboutRepository: donationCategoryRepository,
		DB:              db,
	}
}

func (s *AboutServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.AboutRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.AboutRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToAboutResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *AboutServiceImpl) FindById(ctx context.Context, Id int) (dto.AboutResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.AboutRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.AboutResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.AboutResponse{}, err
		}

		return helper.ToAboutResponse(blocat), nil
	})
	return res.(dto.AboutResponse), err
}

func (s *AboutServiceImpl) Create(ctx context.Context, req *dto.AboutCreate) (dto.AboutResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.AboutResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.About{
		LatarTitle:    req.LatarTitle,
		LatarText:     req.LatarText,
		VisiMisiTitle: req.VisiMisiTitle,
		VisiTitle:     req.VisiTitle,
		VisiText:      req.VisiText,
		MisiTitle:     req.MisiTitle,
		MisiText:      req.MisiText,
		MisiText2:     req.MisiText2,
		NilaiTitle:    req.NilaiTitle,
		NilaiTitle1:   req.NilaiTitle1,
		NilaiText1:    req.NilaiText1,
		NilaiImage1:   req.NilaiImage1,
		NilaiTitle2:   req.NilaiTitle2,
		NilaiText2:    req.NilaiText2,
		NilaiImage2:   req.NilaiImage2,
		NilaiTitle3:   req.NilaiTitle3,
		NilaiText3:    req.NilaiText3,
		NilaiImage3:   req.NilaiImage3,
		NilaiTitle4:   req.NilaiTitle4,
		NilaiText4:    req.NilaiText4,
		NilaiImage4:   req.NilaiImage4,
		StrukturTitle: req.StrukturTitle,
		StrukturImage: req.StrukturImage,
	}

	// save datacategory
	data, err = s.AboutRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.AboutResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.AboutResponse{}, errors.New("error commit transaction")
	}

	return helper.ToAboutResponse(data), nil
}

func (s *AboutServiceImpl) Update(ctx context.Context, req *dto.AboutUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.AboutRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.About{
			Id:            req.Id,
			LatarTitle:    req.LatarTitle,
			LatarText:     req.LatarText,
			VisiMisiTitle: req.VisiMisiTitle,
			VisiTitle:     req.VisiTitle,
			VisiText:      req.VisiText,
			MisiTitle:     req.MisiTitle,
			MisiText:      req.MisiText,
			MisiText2:     req.MisiText2,
			NilaiTitle:    req.NilaiTitle,
			NilaiTitle1:   req.NilaiTitle1,
			NilaiText1:    req.NilaiText1,
			NilaiImage1:   req.NilaiImage1,
			NilaiTitle2:   req.NilaiTitle2,
			NilaiText2:    req.NilaiText2,
			NilaiImage2:   req.NilaiImage2,
			NilaiTitle3:   req.NilaiTitle3,
			NilaiText3:    req.NilaiText3,
			NilaiImage3:   req.NilaiImage3,
			NilaiTitle4:   req.NilaiTitle4,
			NilaiText4:    req.NilaiText4,
			NilaiImage4:   req.NilaiImage4,
			StrukturTitle: req.StrukturTitle,
			StrukturImage: req.StrukturImage,
		}

		// update data category donation
		if err = s.AboutRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *AboutServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.AboutRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.AboutRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
