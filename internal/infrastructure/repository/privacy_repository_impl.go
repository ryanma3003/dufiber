package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type PrivacyRepositoryImpl struct{}

func NewPrivacyRepository() repository.PrivacyRepository {
	return &PrivacyRepositoryImpl{}
}

func (r *PrivacyRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, front *entity.Privacy) (entity.Privacy, error) {
	var id int
	sql := `INSERT INTO privacies (title, text) 
	VALUES ($1, $2) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, front.Title, front.Text)

	if err := result.Scan(&id); err != nil {
		return *front, err
	}

	front.Id = int(id)
	return *front, nil
}

func (r *PrivacyRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, front *entity.Privacy) error {
	sql := `UPDATE privacies SET title=$1, text=$2, updated_at=NOW() WHERE id=$3`
	if _, err := tx.ExecContext(ctx, sql, front.Title, front.Text, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *PrivacyRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, front *entity.Privacy) error {
	sql := `DELETE FROM privacies WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *PrivacyRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Privacy, error) {
	var front entity.Privacy
	sql := `SELECT id, title, text, created_at, updated_at FROM privacies WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.Title, &front.Text, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *PrivacyRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Privacy, error) {
	var fronts []entity.Privacy
	sql := `SELECT id, title, text, created_at, updated_at FROM privacies ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return fronts, err
	}
	defer rows.Close()

	for rows.Next() {
		var front entity.Privacy
		if err := rows.Scan(&front.Id, &front.Title, &front.Text, &front.CreatedAt, &front.UpdatedAt); err != nil {
			return fronts, err
		}
		fronts = append(fronts, front)
	}

	return fronts, nil
}

func (r *PrivacyRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM privacies`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
