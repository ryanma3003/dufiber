package service

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/repository"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type FrontService interface {
	HomepageFirst(ctx context.Context) (dto.HomepageResponse, error)
	AboutFirst(ctx context.Context) (dto.AboutResponse, error)
	TermFirst(ctx context.Context) (dto.TermResponse, error)
	PrivacyFirst(ctx context.Context) (dto.PrivacyResponse, error)
	ContactFirst(ctx context.Context) (dto.ContactResponse, error)
	GaleriAll(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	FaqAll(ctx context.Context) ([]dto.FaqResponse, error)
	BlogAll(ctx context.Context, limit, offset int) (dto.PaginationData, error)
	BlogFindBySlug(ctx context.Context, slug string) (dto.BlogResponse, error)
}

type FrontServiceImpl struct {
	FrontRepository repository.FrontRepository
	DB              *sql.DB
}

func NewFrontService(frontRepository repository.FrontRepository, db *sql.DB) FrontService {
	return &FrontServiceImpl{
		FrontRepository: frontRepository,
		DB:              db,
	}
}

func (s *FrontServiceImpl) HomepageFirst(ctx context.Context) (dto.HomepageResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.FrontRepository.HomepageFindByID(ctx, tx, 1)
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

func (s *FrontServiceImpl) AboutFirst(ctx context.Context) (dto.AboutResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.FrontRepository.AboutFindByID(ctx, tx, 1)
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

func (s *FrontServiceImpl) GaleriAll(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.FrontRepository.GaleriFindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.FrontRepository.GaleriFindTotal(ctx, tx)
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

func (s *FrontServiceImpl) FaqAll(ctx context.Context) ([]dto.FaqResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.FrontRepository.FaqFindAll(ctx, tx)
		if err != nil {
			return []dto.FaqResponse{}, err
		}

		return helper.ToFaqResponses(datas), nil

	})

	return res.([]dto.FaqResponse), err
}

func (s *FrontServiceImpl) TermFirst(ctx context.Context) (dto.TermResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.FrontRepository.TermFindByID(ctx, tx, 1)
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

func (s *FrontServiceImpl) PrivacyFirst(ctx context.Context) (dto.PrivacyResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.FrontRepository.PrivacyFindByID(ctx, tx, 1)
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

func (s *FrontServiceImpl) ContactFirst(ctx context.Context) (dto.ContactResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.FrontRepository.ContactFindByID(ctx, tx, 1)
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

func (s *FrontServiceImpl) BlogAll(ctx context.Context, limit, offset int) (dto.PaginationData, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		datas, err := s.FrontRepository.BlogFindAllWithPagination(ctx, tx, limit, offset)
		if err != nil {
			return dto.PaginationData{}, err
		}

		totalData, err := s.FrontRepository.BlogFindTotal(ctx, tx)
		if err != nil {
			return dto.PaginationData{}, err
		}

		return dto.PaginationData{
			TotalData: totalData,
			Data:      helper.ToBlogResponses(datas),
		}, nil

	})

	return res.(dto.PaginationData), err
}

func (s *FrontServiceImpl) BlogFindBySlug(ctx context.Context, slug string) (dto.BlogResponse, error) {
	res, err := helper.WithTransaction(ctx, s.DB, func(tx *sql.Tx) (interface{}, error) {
		blocat, err := s.FrontRepository.BlogFindBySlug(ctx, tx, slug)
		if err != nil {
			if err == sql.ErrNoRows {
				return dto.BlogResponse{}, helper.NewErrorRowNotFound()
			}
			return dto.BlogResponse{}, err
		}

		return helper.ToBlogResponse(blocat), nil
	})
	return res.(dto.BlogResponse), err
}
