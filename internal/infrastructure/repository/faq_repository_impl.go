package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type FaqRepositoryImpl struct{}

func NewFaqRepository() repository.FaqRepository {
	return &FaqRepositoryImpl{}
}

func (r *FaqRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, front *entity.Faq) (entity.Faq, error) {
	var id int
	sql := `INSERT INTO faq (pertanyaan, jawaban) 
	VALUES ($1, $2) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, front.Pertanyaan, front.Jawaban)

	if err := result.Scan(&id); err != nil {
		return *front, err
	}

	front.Id = int(id)
	return *front, nil
}

func (r *FaqRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, front *entity.Faq) error {
	sql := `UPDATE faq SET pertanyaan=$1, jawaban=$2, updated_at=NOW() WHERE id=$7`
	if _, err := tx.ExecContext(ctx, sql, front.Pertanyaan, front.Jawaban, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *FaqRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, front *entity.Faq) error {
	sql := `DELETE FROM faq WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *FaqRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Faq, error) {
	var front entity.Faq
	sql := `SELECT id, pertanyaan, jawaban, created_at, updated_at FROM faq WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.Pertanyaan, &front.Jawaban, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *FaqRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Faq, error) {
	var fronts []entity.Faq
	sql := `SELECT id, pertanyaan, jawaban, created_at, updated_at FROM faq ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return fronts, err
	}
	defer rows.Close()

	for rows.Next() {
		var front entity.Faq
		if err := rows.Scan(&front.Id, &front.Pertanyaan, &front.Jawaban, &front.CreatedAt, &front.UpdatedAt); err != nil {
			return fronts, err
		}
		fronts = append(fronts, front)
	}

	return fronts, nil
}

func (r *FaqRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM faq`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
