package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type BlogCategoryRepositoryImpl struct{}

func NewBlogCategoryRepository() repository.BlogCategoryRepository {
	return &BlogCategoryRepositoryImpl{}
}

func (r *BlogCategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, blog *entity.BlogCategory) (entity.BlogCategory, error) {
	var id int
	sql := `INSERT INTO blog_category (title, description) VALUES ($1, $2) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, blog.Title, blog.Description)

	if err := result.Scan(&id); err != nil {
		return *blog, err
	}

	blog.Id = int(id)
	return *blog, nil
}

func (r *BlogCategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, blog *entity.BlogCategory) error {
	sql := `UPDATE blog_category SET title=$1, description=$2, updated_at=NOW() WHERE id=$3`
	if _, err := tx.ExecContext(ctx, sql, blog.Title, blog.Description, blog.Id); err != nil {
		return err
	}
	return nil
}

func (r *BlogCategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, blog *entity.BlogCategory) error {
	sql := `DELETE FROM blog_category WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, blog.Id); err != nil {
		return err
	}
	return nil
}

func (r *BlogCategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.BlogCategory, error) {
	var blog entity.BlogCategory
	sql := `SELECT id, title, description, created_at, updated_at FROM blog_category WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&blog.Id, &blog.Title, &blog.Description, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
		return blog, err
	}
	return blog, nil
}

func (r *BlogCategoryRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.BlogCategory, error) {
	var blogs []entity.BlogCategory
	sql := `SELECT id, title, description, created_at, updated_at FROM blog_category ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return blogs, err
	}
	defer rows.Close()

	for rows.Next() {
		var blog entity.BlogCategory
		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Description, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return blogs, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (r *BlogCategoryRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM blog_category`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
