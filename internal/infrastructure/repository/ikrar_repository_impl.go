package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type IkrarRepositoryImpl struct{}

func NewIkrarRepository() repository.IkrarRepository {
	return &IkrarRepositoryImpl{}
}

func (r *IkrarRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, front *entity.Ikrar) (entity.Ikrar, error) {
	var id int
	sql := `INSERT INTO ikrars (nama, email, telepon, tanggal, nama_hari, jumlah_donasi, jumlah_pohon, harga_satu_pohon, nama_pohon) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, front.Nama, front.Email, front.Telepon, front.Tanggal, front.NamaHari, front.JumlahDonasi, front.JumlahPohon, front.HargaSatuPohon, front.NamaPohon)

	if err := result.Scan(&id); err != nil {
		return *front, err
	}

	front.Id = int(id)
	return *front, nil
}

func (r *IkrarRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, front *entity.Ikrar) error {
	sql := `UPDATE ikrars SET nama=$1, email=$2, telepon=$3, tanggal=$4, nama_hari=$5, jumlah_donasi=$6, jumlah_pohon=$7, harga_satu_pohon=$8, nama_pohon=$9, updated_at=NOW() WHERE id=$10`
	if _, err := tx.ExecContext(ctx, sql, front.Nama, front.Email, front.Telepon, front.Tanggal, front.NamaHari, front.JumlahDonasi, front.JumlahPohon, front.HargaSatuPohon, front.NamaPohon, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *IkrarRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, front *entity.Ikrar) error {
	sql := `DELETE FROM ikrars WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *IkrarRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Ikrar, error) {
	var front entity.Ikrar
	sql := `SELECT id, nama, email, telepon, tanggal, nama_hari, jumlah_donasi, jumlah_pohon, harga_satu_pohon, nama_pohon, created_at, updated_at FROM ikrars WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.Nama, &front.Email, &front.Telepon, &front.Tanggal, &front.NamaHari, &front.JumlahDonasi, &front.JumlahPohon, &front.HargaSatuPohon, &front.NamaPohon, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *IkrarRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Ikrar, error) {
	var fronts []entity.Ikrar
	sql := `SELECT id, nama, email, telepon, tanggal, nama_hari, jumlah_donasi, jumlah_pohon, harga_satu_pohon, nama_pohon, created_at, updated_at FROM ikrars ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return fronts, err
	}
	defer rows.Close()

	for rows.Next() {
		var front entity.Ikrar
		if err := rows.Scan(&front.Id, &front.Nama, &front.Email, &front.Telepon, &front.Tanggal, &front.NamaHari, &front.JumlahDonasi, &front.JumlahPohon, &front.HargaSatuPohon, &front.NamaPohon, &front.CreatedAt, &front.UpdatedAt); err != nil {
			return fronts, err
		}
		fronts = append(fronts, front)
	}

	return fronts, nil
}

func (r *IkrarRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM ikrars`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
