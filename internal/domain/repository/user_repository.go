package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user *entity.User) (entity.User, error)
	Update(ctx context.Context, tx *sql.Tx, user *entity.User) error
	Delete(ctx context.Context, tx *sql.Tx, user *entity.User) error
	ChangePassword(ctx context.Context, tx *sql.Tx, user *entity.User) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.User, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}
