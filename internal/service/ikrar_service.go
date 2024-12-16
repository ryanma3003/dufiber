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

type IkrarService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.IkrarResponse, error)
	Create(ctx context.Context, req *dto.IkrarCreate) (dto.IkrarResponse, error)
	Update(ctx context.Context, req *dto.IkrarUpdate) error
	Delete(ctx context.Context, id int) error
}

type IkrarServiceImpl struct {
	IkrarRepository repository.IkrarRepository
	DB              *sql.DB
}

func NewIkrarService(donationCategoryRepository repository.IkrarRepository, db *sql.DB) IkrarService {
	return &IkrarServiceImpl{
		IkrarRepository: donationCategoryRepository,
		DB:              db,
	}
}

func (s *IkrarServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.IkrarRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.IkrarRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToIkrarResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *IkrarServiceImpl) FindById(ctx context.Context, Id int) (dto.IkrarResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.IkrarRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.IkrarResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.IkrarResponse{}, err
		}

		return helper.ToIkrarResponse(blocat), nil
	})
	return res.(dto.IkrarResponse), err
}

func (s *IkrarServiceImpl) Create(ctx context.Context, req *dto.IkrarCreate) (dto.IkrarResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.IkrarResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.Ikrar{
		Nama:           req.Nama,
		Email:          req.Email,
		Telepon:        req.Telepon,
		Tanggal:        req.Tanggal,
		NamaHari:       req.NamaHari,
		JumlahDonasi:   req.JumlahDonasi,
		JumlahPohon:    req.JumlahPohon,
		HargaSatuPohon: req.HargaSatuPohon,
		NamaPohon:      req.NamaPohon,
	}

	// save datacategory
	data, err = s.IkrarRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.IkrarResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.IkrarResponse{}, errors.New("error commit transaction")
	}

	return helper.ToIkrarResponse(data), nil
}

func (s *IkrarServiceImpl) Update(ctx context.Context, req *dto.IkrarUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.IkrarRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.Ikrar{
			Id:             req.Id,
			Nama:           req.Nama,
			Email:          req.Email,
			Telepon:        req.Telepon,
			Tanggal:        req.Tanggal,
			NamaHari:       req.NamaHari,
			JumlahDonasi:   req.JumlahDonasi,
			JumlahPohon:    req.JumlahPohon,
			HargaSatuPohon: req.HargaSatuPohon,
			NamaPohon:      req.NamaPohon,
		}

		// update data category donation
		if err = s.IkrarRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *IkrarServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.IkrarRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.IkrarRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
