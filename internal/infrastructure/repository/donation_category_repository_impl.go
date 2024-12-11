package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type DonationCategoryRepositoryImpl struct{}

func NewDonationCategoryRepository() repository.DonationCategoryRepository {
	return &DonationCategoryRepositoryImpl{}
}

func (r *DonationCategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, donate *entity.DonationCategory) (entity.DonationCategory, error) {
	var id int
	sql := `INSERT INTO donation_category (title, description) VALUES ($1, $2) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, donate.Title, donate.Description)

	if err := result.Scan(&id); err != nil {
		return *donate, err
	}

	donate.Id = int(id)
	return *donate, nil
}

func (r *DonationCategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, donate *entity.DonationCategory) error {
	sql := `UPDATE donation_category SET title=$1, description=$2, updated_at=NOW() WHERE id=$3`
	if _, err := tx.ExecContext(ctx, sql, donate.Title, donate.Description, donate.Id); err != nil {
		return err
	}
	return nil
}

func (r *DonationCategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, donate *entity.DonationCategory) error {
	sql := `DELETE FROM donation_category WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, donate.Id); err != nil {
		return err
	}
	return nil
}

func (r *DonationCategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.DonationCategory, error) {
	var donate entity.DonationCategory
	sql := `SELECT id, title, description, created_at, updated_at FROM donation_category WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&donate.Id, &donate.Title, &donate.Description, &donate.CreatedAt, &donate.UpdatedAt); err != nil {
		return donate, err
	}
	return donate, nil
}

func (r *DonationCategoryRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.DonationCategory, error) {
	var donates []entity.DonationCategory
	sql := `SELECT id, title, description, created_at, updated_at FROM donation_category ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return donates, err
	}
	defer rows.Close()

	for rows.Next() {
		var donate entity.DonationCategory
		if err := rows.Scan(&donate.Id, &donate.Title, &donate.Description, &donate.CreatedAt, &donate.UpdatedAt); err != nil {
			return donates, err
		}
		donates = append(donates, donate)
	}

	return donates, nil
}

func (r *DonationCategoryRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM donation_category`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
