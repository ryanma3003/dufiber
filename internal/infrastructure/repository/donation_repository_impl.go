package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type DonationRepositoryImpl struct{}

func NewDonationRepository() repository.DonationRepository {
	return &DonationRepositoryImpl{}
}

func (r *DonationRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, donate *entity.Donation) (entity.Donation, error) {
	var id int
	sql := `INSERT INTO donations (name, email, phone, amount, status, reference, snap_token, donation_list_id, charity_list_id, user_id, orderId, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW()) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, donate.Name, donate.Email, donate.Phone, donate.Amount, donate.Status, donate.Reference, donate.SnapToken, donate.DonationListId, donate.CharityListId, donate.UserId, donate.OrderId)

	if err := result.Scan(&id); err != nil {
		return *donate, err
	}

	donate.Id = int(id)
	return *donate, nil
}

func (r *DonationRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, donate *entity.Donation) error {
	sql := `UPDATE donations SET name=?, email=?, phone=?, amount=?, status=?, reference=?, snap_token=?, donation_list_id=?, charity_list_id=?, user_id=?, orderId=?, updated_at=NOW() WHERE id=?`
	if _, err := tx.ExecContext(ctx, sql, donate.Name, donate.Email, donate.Phone, donate.Status, donate.Reference, donate.SnapToken, donate.DonationListId, donate.CharityListId, donate.UserId, donate.OrderId, donate.Id); err != nil {
		return err
	}
	return nil
}

func (r *DonationRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, donate *entity.Donation) error {
	sql := `DELETE FROM donations WHERE id=?`
	if _, err := tx.ExecContext(ctx, sql, donate.Id); err != nil {
		return err
	}
	return nil
}

func (r *DonationRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Donation, error) {
	var donate entity.Donation
	sql := `SELECT id, name, email, phone, amount, status, reference, snap_token, donation_list_id, charity_list_id, user_id, orderId, created_at, updated_at FROM donations WHERE id=?`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&donate.Id, &donate.Name, &donate.Email, &donate.Phone, &donate.Amount, &donate.Status, &donate.Reference, &donate.SnapToken, &donate.DonationListId, &donate.CharityListId, &donate.UserId, donate.OrderId, &donate.CreatedAt, &donate.UpdatedAt); err != nil {
		return donate, err
	}
	return donate, nil
}

func (r *DonationRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Donation, error) {
	var donates []entity.Donation
	sql := `SELECT id, name, email, phone, amount, status, reference, snap_token, donation_list_id, charity_list_id, user_id, orderId, created_at, updated_at FROM donations ORDER BY created_at DESC LIMIT ? OFFSET ?`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return donates, err
	}
	defer rows.Close()

	for rows.Next() {
		var donate entity.Donation
		if err := rows.Scan(&donate.Id, &donate.Name, &donate.Email, &donate.Phone, &donate.Amount, &donate.Status, &donate.Reference, &donate.SnapToken, &donate.DonationListId, &donate.CharityListId, &donate.UserId, donate.OrderId, &donate.CreatedAt, &donate.UpdatedAt); err != nil {
			return donates, err
		}
		donates = append(donates, donate)
	}

	return donates, nil
}

func (r *DonationRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM donations`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}

func (r *DonationRepositoryImpl) FindTotalDonatur(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) AS total FROM donations WHERE status = 'success'`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}

func (r *DonationRepositoryImpl) FindTotalZakat(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COALESCE(SUM(amount),0) AS total FROM donations WHERE donation_list_id IN(1,2,3,4,5,6) AND status = 'success'`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}

func (r *DonationRepositoryImpl) FindTotalInfaq(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COALESCE(SUM(amount),0) AS total FROM donations WHERE donation_list_id IN(7,8,9,10,11,12) AND status = 'success'`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}

func (r *DonationRepositoryImpl) FindTotalWakaf(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COALESCE(SUM(amount),0) AS total FROM donations WHERE donation_list_id = 13 AND status = 'success'`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
