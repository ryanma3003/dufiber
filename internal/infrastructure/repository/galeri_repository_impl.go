package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type GaleriRepositoryImpl struct{}

func NewGaleriRepository() repository.GaleriRepository {
	return &GaleriRepositoryImpl{}
}

func (r *GaleriRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, galeri *entity.Galeri) (entity.Galeri, error) {
	var id int
	sql := `INSERT INTO galeris (title, slug, image, galeri_tag_id) VALUES ($1, $2, $3, $4) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, galeri.Title, galeri.Slug, galeri.Image, galeri.GaleriTagId)

	if err := result.Scan(&id); err != nil {
		return *galeri, err
	}

	galeri.Id = int(id)
	return *galeri, nil
}

func (r *GaleriRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, galeri *entity.Galeri) error {
	sql := `UPDATE galeris SET title=$1, slug=$2, image=$3, galeri_tag_id=$4, updated_at=NOW() WHERE id=$5`
	if _, err := tx.ExecContext(ctx, sql, galeri.Title, galeri.Slug, galeri.Image, galeri.GaleriTagId, galeri.Id); err != nil {
		return err
	}
	return nil
}

func (r *GaleriRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, galeri *entity.Galeri) error {
	sql := `DELETE FROM galeris WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, galeri.Id); err != nil {
		return err
	}
	return nil
}

func (r *GaleriRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Galeri, error) {
	var galeri entity.Galeri
	sql := `SELECT id, title, slug, image, galeri_tag_id, created_at, updated_at FROM galeris WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&galeri.Id, &galeri.Title, &galeri.Slug, galeri.Image, galeri.GaleriTagId, &galeri.CreatedAt, &galeri.UpdatedAt); err != nil {
		return galeri, err
	}
	return galeri, nil
}

func (r *GaleriRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Galeri, error) {
	var galeris []entity.Galeri
	sql := `SELECT id, title, slug, image, galeri_tag_id, created_at, updated_at FROM galeris ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return galeris, err
	}
	defer rows.Close()

	for rows.Next() {
		var galeri entity.Galeri
		if err := rows.Scan(&galeri.Id, &galeri.Title, &galeri.Slug, galeri.Image, galeri.GaleriTagId, &galeri.CreatedAt, &galeri.UpdatedAt); err != nil {
			return galeris, err
		}
		galeris = append(galeris, galeri)
	}

	return galeris, nil
}

func (r *GaleriRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM galeris`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
