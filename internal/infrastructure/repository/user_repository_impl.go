package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type UserRepositoryImpl struct{}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user *entity.User) (entity.User, error) {
	var id int
	sql := `INSERT INTO users (username, password, email, role) VALUES ($1, $2, $3, $4) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, user.Username, user.Password, user.Email, user.Role)

	if err := result.Scan(&id); err != nil {
		return *user, err
	}

	user.Id = int(id)
	return *user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *entity.User) error {
	sql := `UPDATE users SET username=$1, email=$2, role=$3, updated_at=NOW() WHERE id=$4`
	if _, err := tx.ExecContext(ctx, sql, user.Username, user.Email, user.Role, user.Id); err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user *entity.User) error {
	sql := `UPDATE users SET is_deleted=true, deleted_at=NOW() WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, user.Id); err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) ChangePassword(ctx context.Context, tx *sql.Tx, user *entity.User) error {
	sql := `UPDATE users SET password=$1, updated_at=NOW() WHERE id=$2`
	if _, err := tx.ExecContext(ctx, sql, user.Password, user.Id); err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.User, error) {
	var user entity.User
	sql := `SELECT id, username, email, role, created_at, updated_at FROM users WHERE id=$1 AND is_deleted=false`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error) {
	var user entity.User
	sql := `SELECT id, username, email, role, created_at, updated_at FROM users WHERE username=$1 AND is_deleted=false`
	if err := tx.QueryRowContext(ctx, sql, username).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.User, error) {
	var users []entity.User
	sql := `SELECT id, username, email, role, created_at, updated_at FROM users WHERE is_deleted=false ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM users WHERE is_deleted=false`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
