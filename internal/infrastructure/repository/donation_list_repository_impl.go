package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type DonationListRepositoryImpl struct{}

func NewDonationListRepository() repository.DonationListRepository {
	return &DonationListRepositoryImpl{}
}

func (r *DonationListRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, donate *entity.DonationList) (entity.DonationList, error) {
	var id int
	sql := `INSERT INTO donation_list (title, description, code, donation_category_id) VALUES ($1, $2, $3, $4) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, donate.Title, donate.Description, donate.Code, donate.DonationCategoryId)

	if err := result.Scan(&id); err != nil {
		return *donate, err
	}

	donate.Id = int(id)
	return *donate, nil
}

func (r *DonationListRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, donate *entity.DonationList) error {
	sql := `UPDATE donation_list SET title=$1, description=$2, code=$3, donation_category_id=$4, updated_at=NOW() WHERE id=$5`
	if _, err := tx.ExecContext(ctx, sql, donate.Title, donate.Description, donate.Code, donate.DonationCategoryId, donate.Id); err != nil {
		return err
	}
	return nil
}

func (r *DonationListRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, donate *entity.DonationList) error {
	sql := `DELETE FROM donation_list WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, donate.Id); err != nil {
		return err
	}
	return nil
}

func (r *DonationListRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.DonationList, error) {
	var donate entity.DonationList
	sql := `SELECT id, title, description, code, donation_category_id, created_at, updated_at FROM donation_list WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&donate.Id, &donate.Title, &donate.Description, &donate.Code, &donate.DonationCategoryId, &donate.CreatedAt, &donate.UpdatedAt); err != nil {
		return donate, err
	}
	return donate, nil
}

func (r *DonationListRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.DonationList, error) {
	var donates []entity.DonationList
	sql := `SELECT id, title, description, code, donation_category_id, created_at, updated_at FROM donation_list ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return donates, err
	}
	defer rows.Close()

	for rows.Next() {
		var donate entity.DonationList
		if err := rows.Scan(&donate.Id, &donate.Title, &donate.Description, &donate.Code, &donate.DonationCategoryId, &donate.CreatedAt, &donate.UpdatedAt); err != nil {
			return donates, err
		}
		donates = append(donates, donate)
	}

	return donates, nil
}

func (r *DonationListRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM donation_list`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
