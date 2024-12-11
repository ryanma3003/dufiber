package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
)

type DonationCategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, doncat *entity.DonationCategory) (entity.DonationCategory, error)
	Update(ctx context.Context, tx *sql.Tx, doncat *entity.DonationCategory) error
	Delete(ctx context.Context, tx *sql.Tx, doncat *entity.DonationCategory) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.DonationCategory, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.DonationCategory, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type DonationListRepository interface {
	Save(ctx context.Context, tx *sql.Tx, donlist *entity.DonationList) (entity.DonationList, error)
	Update(ctx context.Context, tx *sql.Tx, donlist *entity.DonationList) error
	Delete(ctx context.Context, tx *sql.Tx, donlist *entity.DonationList) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.DonationList, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.DonationList, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type DonationRepository interface {
	Save(ctx context.Context, tx *sql.Tx, donation *entity.Donation) (entity.Donation, error)
	Update(ctx context.Context, tx *sql.Tx, donation *entity.Donation) error
	Delete(ctx context.Context, tx *sql.Tx, donation *entity.Donation) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Donation, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Donation, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type HargaZakatRepository interface {
	Save(ctx context.Context, tx *sql.Tx, hargazakat *entity.HargaZakat) (entity.HargaZakat, error)
	Update(ctx context.Context, tx *sql.Tx, hargazakat *entity.HargaZakat) error
	Delete(ctx context.Context, tx *sql.Tx, hargazakat *entity.HargaZakat) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.HargaZakat, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.HargaZakat, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}
