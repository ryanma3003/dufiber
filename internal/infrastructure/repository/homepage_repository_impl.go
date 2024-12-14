package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type HomepageRepositoryImpl struct{}

func NewHomepageRepository() repository.HomepageRepository {
	return &HomepageRepositoryImpl{}
}

func (r *HomepageRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, front *entity.Homepage) (entity.Homepage, error) {
	var id int
	sql := `INSERT INTO homepages (main_image, main_title, main_text, kalkulator_title, kalulator_text, pers_text, publikasi_text) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, front.MainImage, front.MainTitle, front.MainText, front.KalkulatorTitle, front.KalkulatorText, front.PersText, front.PublikasiText)

	if err := result.Scan(&id); err != nil {
		return *front, err
	}

	front.Id = int(id)
	return *front, nil
}

func (r *HomepageRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, front *entity.Homepage) error {
	sql := `UPDATE homepages SET main_image=$1, main_title=$2, main_text=$3, kalkulator_title=$4, kalkulator_text=$5, pers_text=$6, publikasi_text=$7, updated_at=NOW() WHERE id=$8`
	if _, err := tx.ExecContext(ctx, sql, front.MainImage, front.MainTitle, front.MainText, front.KalkulatorTitle, front.KalkulatorText, front.PersText, front.PublikasiText, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *HomepageRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, front *entity.Homepage) error {
	sql := `DELETE FROM homepages WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *HomepageRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Homepage, error) {
	var front entity.Homepage
	sql := `SELECT id, main_image, main_title, main_text, kalkulator_title, kalkulator_text, pers_text, publikasi_text, created_at, updated_at FROM homepages WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.MainImage, &front.MainTitle, &front.MainText, &front.KalkulatorTitle, &front.KalkulatorText, &front.PersText, &front.PublikasiText, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *HomepageRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Homepage, error) {
	var fronts []entity.Homepage
	sql := `SELECT id, main_image, main_title, main_text, kalkulator_title, kalkulator_text, pers_text, publikasi_text, created_at, updated_at FROM homepages ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return fronts, err
	}
	defer rows.Close()

	for rows.Next() {
		var front entity.Homepage
		if err := rows.Scan(&front.Id, &front.MainImage, &front.MainTitle, &front.MainText, &front.KalkulatorTitle, &front.KalkulatorText, &front.PersText, &front.PublikasiText, &front.CreatedAt, &front.UpdatedAt); err != nil {
			return fronts, err
		}
		fronts = append(fronts, front)
	}

	return fronts, nil
}

func (r *HomepageRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM homepages`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
