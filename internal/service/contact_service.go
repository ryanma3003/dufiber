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

type ContactService interface {
	FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FindById(ctx context.Context, id int) (dto.ContactResponse, error)
	Create(ctx context.Context, req *dto.ContactCreate) (dto.ContactResponse, error)
	Update(ctx context.Context, req *dto.ContactUpdate) error
	Delete(ctx context.Context, id int) error
}

type ContactServiceImpl struct {
	ContactRepository repository.ContactRepository
	DB                *sql.DB
}

func NewContactService(donationCategoryRepository repository.ContactRepository, db *sql.DB) ContactService {
	return &ContactServiceImpl{
		ContactRepository: donationCategoryRepository,
		DB:                db,
	}
}

func (s *ContactServiceImpl) FindAllWithPagination(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.ContactRepository.FindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.ContactRepository.FindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToContactResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *ContactServiceImpl) FindById(ctx context.Context, Id int) (dto.ContactResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.ContactRepository.FindByID(ctx, tx, Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.ContactResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.ContactResponse{}, err
		}

		return helper.ToContactResponse(blocat), nil
	})
	return res.(dto.ContactResponse), err
}

func (s *ContactServiceImpl) Create(ctx context.Context, req *dto.ContactCreate) (dto.ContactResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return dto.ContactResponse{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("rollback error: %v (original error: %w)", rbErr, err)
			}
		}
	}()

	data := entity.Contact{
		MainTag:      req.MainTag,
		MainText:     req.MainText,
		AddressTitle: req.AddressTitle,
		Address:      req.Address,
		PhoneTitle:   req.PhoneTitle,
		Phone:        req.Phone,
	}

	// save datacategory
	data, err = s.ContactRepository.Save(ctx, tx, &data)
	if err != nil {
		return dto.ContactResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return dto.ContactResponse{}, errors.New("error commit transaction")
	}

	return helper.ToContactResponse(data), nil
}

func (s *ContactServiceImpl) Update(ctx context.Context, req *dto.ContactUpdate) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {

		// check data category by id
		_, err := s.ContactRepository.FindByID(ctx, tx, req.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		data_update := entity.Contact{
			Id:           req.Id,
			MainTag:      req.MainTag,
			MainText:     req.MainText,
			AddressTitle: req.AddressTitle,
			Address:      req.Address,
			PhoneTitle:   req.PhoneTitle,
			Phone:        req.Phone,
		}

		// update data category donation
		if err = s.ContactRepository.Update(ctx, tx, &data_update); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func (s *ContactServiceImpl) Delete(ctx context.Context, id int) error {
	_, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		// check data category by id
		datacat, err := s.ContactRepository.FindByID(ctx, tx, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, helper.NewErrorRowNotFound()
			}

			return nil, err
		}

		// delete data category
		if err = s.ContactRepository.Delete(ctx, tx, &datacat); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}
