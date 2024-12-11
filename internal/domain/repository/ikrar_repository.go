package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
)

type IkrarRepository interface {
	Save(ctx context.Context, tx *sql.Tx, ikrar *entity.Ikrar) (entity.Ikrar, error)
	Update(ctx context.Context, tx *sql.Tx, ikrar *entity.Ikrar) error
	Delete(ctx context.Context, tx *sql.Tx, ikrar *entity.Ikrar) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Ikrar, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Ikrar, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}
