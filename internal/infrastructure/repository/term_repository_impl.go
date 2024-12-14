package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type TermRepositoryImpl struct{}

func NewTermRepository() repository.TermRepository {
	return &TermRepositoryImpl{}
}

func (r *TermRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, front *entity.Term) (entity.Term, error) {
	var id int
	sql := `INSERT INTO terms (title, text) 
	VALUES ($1, $2) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, front.Title, front.Text)

	if err := result.Scan(&id); err != nil {
		return *front, err
	}

	front.Id = int(id)
	return *front, nil
}

func (r *TermRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, front *entity.Term) error {
	sql := `UPDATE terms SET title=$1, text=$2, updated_at=NOW() WHERE id=$3`
	if _, err := tx.ExecContext(ctx, sql, front.Title, front.Text, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *TermRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, front *entity.Term) error {
	sql := `DELETE FROM terms WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *TermRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Term, error) {
	var front entity.Term
	sql := `SELECT id, title, text, created_at, updated_at FROM terms WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.Title, &front.Text, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *TermRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Term, error) {
	var fronts []entity.Term
	sql := `SELECT id, title, text, created_at, updated_at FROM terms ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return fronts, err
	}
	defer rows.Close()

	for rows.Next() {
		var front entity.Term
		if err := rows.Scan(&front.Id, &front.Title, &front.Text, &front.CreatedAt, &front.UpdatedAt); err != nil {
			return fronts, err
		}
		fronts = append(fronts, front)
	}

	return fronts, nil
}

func (r *TermRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM terms`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
