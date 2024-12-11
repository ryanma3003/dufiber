package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
)

type HomepageRepository interface {
	Save(ctx context.Context, tx *sql.Tx, homepage *entity.Homepage) (entity.Homepage, error)
	Update(ctx context.Context, tx *sql.Tx, homepage *entity.Homepage) error
	Delete(ctx context.Context, tx *sql.Tx, homepage *entity.Homepage) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Homepage, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Homepage, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type FaqRepository interface {
	Save(ctx context.Context, tx *sql.Tx, faq *entity.Faq) (entity.Faq, error)
	Update(ctx context.Context, tx *sql.Tx, faq *entity.Faq) error
	Delete(ctx context.Context, tx *sql.Tx, faq *entity.Faq) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Faq, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Faq, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type TermRepository interface {
	Save(ctx context.Context, tx *sql.Tx, term *entity.Term) (entity.Term, error)
	Update(ctx context.Context, tx *sql.Tx, term *entity.Term) error
	Delete(ctx context.Context, tx *sql.Tx, term *entity.Term) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Term, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Term, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type PrivacyRepository interface {
	Save(ctx context.Context, tx *sql.Tx, privacy *entity.Privacy) (entity.Privacy, error)
	Update(ctx context.Context, tx *sql.Tx, privacy *entity.Privacy) error
	Delete(ctx context.Context, tx *sql.Tx, privacy *entity.Privacy) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Privacy, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Privacy, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}

type ContactRepository interface {
	Save(ctx context.Context, tx *sql.Tx, contact *entity.Contact) (entity.Contact, error)
	Update(ctx context.Context, tx *sql.Tx, contact *entity.Contact) error
	Delete(ctx context.Context, tx *sql.Tx, contact *entity.Contact) error
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Contact, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Contact, error)
	FindTotal(ctx context.Context, tx *sql.Tx) (int, error)
}
