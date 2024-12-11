package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
)

type GaleriTagRepository interface {
	Save(ctx context.Context, tx *sql.Tx, galeritag *entity.GaleriTag) (entity.GaleriTag, error)
	Update(ctx context.Context, tx *sql.Tx, galeritag *entity.GaleriTag) error
	Delete(ctx context.Context, tx *sql.Tx, galeritag *entity.GaleriTag) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.GaleriTag, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.GaleriTag, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type GaleriRepository interface {
	Save(ctx context.Context, tx *sql.Tx, galeri *entity.Galeri) (entity.Galeri, error)
	Update(ctx context.Context, tx *sql.Tx, galeri *entity.Galeri) error
	Delete(ctx context.Context, tx *sql.Tx, galeri *entity.Galeri) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Galeri, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Galeri, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}
