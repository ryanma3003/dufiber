package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type BlogRepositoryImpl struct{}

func NewBlogRepository() repository.BlogRepository {
	return &BlogRepositoryImpl{}
}

func (r *BlogRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, blog *entity.Blog) (entity.Blog, error) {
	var id int
	sql := `INSERT INTO blog (title, image, content, slug, author, blog_category_id, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, blog.Title, blog.Image, blog.Content, blog.Slug, blog.Author, blog.BlogCategoryId, blog.UserId)

	if err := result.Scan(&id); err != nil {
		return *blog, err
	}

	blog.Id = int(id)
	return *blog, nil
}

func (r *BlogRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, blog *entity.Blog) error {
	sql := `UPDATE blog SET title=$1, content=$2, image=$3, slug=$4, author=$5, blog_category_id=$6, user_id=$7, updated_at=NOW() WHERE id=$8`
	if _, err := tx.ExecContext(ctx, sql, blog.Title, blog.Content, blog.Image, blog.Slug, blog.Author, blog.BlogCategoryId, blog.UserId, blog.Id); err != nil {
		return err
	}
	return nil
}

func (r *BlogRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, blog *entity.Blog) error {
	sql := `DELETE FROM blog WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, blog.Id); err != nil {
		return err
	}
	return nil
}

func (r *BlogRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Blog, error) {
	var blog entity.Blog
	sql := `SELECT id, title, content, image, author, user_id, blog_category_id, slug, created_at, updated_at FROM blog WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Image, &blog.Author, &blog.UserId, &blog.BlogCategoryId, &blog.Slug, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
		return blog, err
	}
	return blog, nil
}

func (r *BlogRepositoryImpl) FindBySlug(ctx context.Context, tx *sql.Tx, slug string) (entity.Blog, error) {
	var blog entity.Blog
	sql := `SELECT id, title, image, slug, content, author, blog_category_id, user_id, created_at, updated_at FROM blog WHERE slug=$1`
	if err := tx.QueryRowContext(ctx, sql, slug).Scan(&blog.Id, &blog.Title, &blog.Image, &blog.Slug, &blog.Content, &blog.Author, &blog.BlogCategoryId, &blog.UserId, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
		return blog, err
	}
	return blog, nil
}

func (r *BlogRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Blog, error) {
	var blogs []entity.Blog
	sql := `SELECT id, title, image, content, slug, author, blog_category_id, user_id, created_at, updated_at FROM blog ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return blogs, err
	}
	defer rows.Close()

	for rows.Next() {
		var blog entity.Blog
		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Image, &blog.Slug, &blog.BlogCategoryId, &blog.Author, &blog.UserId, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return blogs, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (r *BlogRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM blog`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
