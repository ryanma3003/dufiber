package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
)

type BlogRepository interface {
	Save(ctx context.Context, tx *sql.Tx, blog *entity.Blog) (entity.Blog, error)
	Update(ctx context.Context, tx *sql.Tx, blog *entity.Blog) error
	Delete(ctx context.Context, tx *sql.Tx, blog *entity.Blog) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Blog, error)
	FindBySlug(ctx context.Context, tx *sql.Tx, slug string) (entity.Blog, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Blog, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type BlogCategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, blogcat *entity.BlogCategory) (entity.BlogCategory, error)
	Update(ctx context.Context, tx *sql.Tx, blogcat *entity.BlogCategory) error
	Delete(ctx context.Context, tx *sql.Tx, blogcat *entity.BlogCategory) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.BlogCategory, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.BlogCategory, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}
