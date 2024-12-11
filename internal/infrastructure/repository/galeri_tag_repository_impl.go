package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type GaleriTagRepositoryImpl struct{}

func NewGaleriTagRepository() repository.GaleriTagRepository {
	return &GaleriTagRepositoryImpl{}
}

func (r *GaleriTagRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, galeri *entity.GaleriTag) (entity.GaleriTag, error) {
	var id int
	sql := `INSERT INTO galeri_tags (title, slug) VALUES ($1, $2) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, galeri.Title, galeri.Slug)

	if err := result.Scan(&id); err != nil {
		return *galeri, err
	}

	galeri.Id = int(id)
	return *galeri, nil
}

func (r *GaleriTagRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, galeri *entity.GaleriTag) error {
	sql := `UPDATE galeri_tags SET title=$1, slug=$2, updated_at=NOW() WHERE id=$3`
	if _, err := tx.ExecContext(ctx, sql, galeri.Title, galeri.Slug, galeri.Id); err != nil {
		return err
	}
	return nil
}

func (r *GaleriTagRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, galeri *entity.GaleriTag) error {
	sql := `DELETE FROM galeri_tags WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, galeri.Id); err != nil {
		return err
	}
	return nil
}

func (r *GaleriTagRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.GaleriTag, error) {
	var galeri entity.GaleriTag
	sql := `SELECT id, title, slug, created_at, updated_at FROM galeri_tags WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&galeri.Id, &galeri.Title, &galeri.Slug, &galeri.CreatedAt, &galeri.UpdatedAt); err != nil {
		return galeri, err
	}
	return galeri, nil
}

func (r *GaleriTagRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.GaleriTag, error) {
	var galeris []entity.GaleriTag
	sql := `SELECT id, title, slug, created_at, updated_at FROM galeri_tags ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return galeris, err
	}
	defer rows.Close()

	for rows.Next() {
		var galeri entity.GaleriTag
		if err := rows.Scan(&galeri.Id, &galeri.Title, &galeri.Slug, &galeri.CreatedAt, &galeri.UpdatedAt); err != nil {
			return galeris, err
		}
		galeris = append(galeris, galeri)
	}

	return galeris, nil
}

func (r *GaleriTagRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM galeri_tags`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
