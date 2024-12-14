package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type ContactRepositoryImpl struct{}

func NewContactRepository() repository.ContactRepository {
	return &ContactRepositoryImpl{}
}

func (r *ContactRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, front *entity.Contact) (entity.Contact, error) {
	var id int
	sql := `INSERT INTO contacts (main_tag, main_text, address_title, address, phone_title, phone) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, front.MainTag, front.MainText, front.AddressTitle, front.Address, front.PhoneTitle, front.Phone)

	if err := result.Scan(&id); err != nil {
		return *front, err
	}

	front.Id = int(id)
	return *front, nil
}

func (r *ContactRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, front *entity.Contact) error {
	sql := `UPDATE contacts SET main_tag=$1, main_text=$2, address_title=$3, address=$4, phone_title=$5, phone=$6, updated_at=NOW() WHERE id=$7`
	if _, err := tx.ExecContext(ctx, sql, front.MainTag, front.MainText, front.AddressTitle, front.Address, front.PhoneTitle, front.Phone, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, front *entity.Contact) error {
	sql := `DELETE FROM contacts WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Contact, error) {
	var front entity.Contact
	sql := `SELECT id, main_tag, main_text, address_title, address, phone_title, phone, created_at, updated_at FROM contacts WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.MainTag, &front.MainText, &front.AddressTitle, &front.Address, &front.PhoneTitle, &front.Phone, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *ContactRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Contact, error) {
	var fronts []entity.Contact
	sql := `SELECT id, main_tag, main_text, address_title, address, phone_title, phone, created_at, updated_at FROM contacts ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return fronts, err
	}
	defer rows.Close()

	for rows.Next() {
		var front entity.Contact
		if err := rows.Scan(&front.Id, &front.MainTag, &front.MainText, &front.AddressTitle, &front.Address, &front.PhoneTitle, &front.Phone, &front.CreatedAt, &front.UpdatedAt); err != nil {
			return fronts, err
		}
		fronts = append(fronts, front)
	}

	return fronts, nil
}

func (r *ContactRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM contacts`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
