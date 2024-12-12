package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type HargaZakatRepositoryImpl struct{}

func NewHargaZakatRepository() repository.HargaZakatRepository {
	return &HargaZakatRepositoryImpl{}
}

func (r *HargaZakatRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, donate *entity.HargaZakat) (entity.HargaZakat, error) {
	var id int
	sql := `INSERT INTO harga_zakats (title, price, donation_list_id) VALUES ($1, $2, $3) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, donate.Title, donate.Price, donate.DonationListId)

	if err := result.Scan(&id); err != nil {
		return *donate, err
	}

	donate.Id = int(id)
	return *donate, nil
}

func (r *HargaZakatRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, donate *entity.HargaZakat) error {
	sql := `UPDATE harga_zakats SET title=$1, price=$2, donation_list_id=$3, updated_at=NOW() WHERE id=$4`
	if _, err := tx.ExecContext(ctx, sql, donate.Title, donate.Price, donate.DonationListId, donate.Id); err != nil {
		return err
	}
	return nil
}

func (r *HargaZakatRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, donate *entity.HargaZakat) error {
	sql := `DELETE FROM harga_zakats WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, donate.Id); err != nil {
		return err
	}
	return nil
}

func (r *HargaZakatRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.HargaZakat, error) {
	var donate entity.HargaZakat
	sql := `SELECT id, title, price, donation_list_id, created_at, updated_at FROM harga_zakats WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&donate.Id, &donate.Title, &donate.Price, &donate.DonationListId, &donate.CreatedAt, &donate.UpdatedAt); err != nil {
		return donate, err
	}
	return donate, nil
}

func (r *HargaZakatRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.HargaZakat, error) {
	var donates []entity.HargaZakat
	sql := `SELECT id, title, price, donation_list_id, created_at, updated_at FROM harga_zakats ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return donates, err
	}
	defer rows.Close()

	for rows.Next() {
		var donate entity.HargaZakat
		if err := rows.Scan(&donate.Id, &donate.Title, &donate.Price, &donate.DonationListId, &donate.CreatedAt, &donate.UpdatedAt); err != nil {
			return donates, err
		}
		donates = append(donates, donate)
	}

	return donates, nil
}

func (r *HargaZakatRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM harga_zakats`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
