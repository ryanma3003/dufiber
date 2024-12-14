package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type AboutRepositoryImpl struct{}

func NewAboutRepository() repository.AboutRepository {
	return &AboutRepositoryImpl{}
}

func (r *AboutRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, front *entity.About) (entity.About, error) {
	var id int
	sql := `INSERT INTO abouts (latar_title, latar_text, visi_misi_title, visi_title, visi_text, misi_title, misi_text, misi_text2, nilai_title, nilai_title1, nilai_text1, nilai_image1, nilai_title2, nilai_text2, nilai_image2, nilai_title3, nilai_text3, nilai_image3, nilai_title4, nilai_text4, nilai_image4, struktur_title, struktur_image) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23) RETURNING id`
	result := tx.QueryRowContext(ctx, sql, front.LatarTitle, front.LatarText, front.VisiMisiTitle, front.VisiTitle, front.VisiText, front.MisiTitle, front.MisiText, front.MisiText2, front.NilaiTitle, front.NilaiTitle1, front.NilaiText1, front.NilaiImage1, front.NilaiTitle2, front.NilaiText2, front.NilaiImage2, front.NilaiTitle3, front.NilaiText3, front.NilaiImage3, front.NilaiTitle4, front.NilaiText4, front.NilaiImage4, front.StrukturTitle, front.StrukturImage)

	if err := result.Scan(&id); err != nil {
		return *front, err
	}

	front.Id = int(id)
	return *front, nil
}

func (r *AboutRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, front *entity.About) error {
	sql := `UPDATE abouts SET latar_title=$1, latar_text=$2, visi_misi_title=$3, visi_title=$4, visi_text=$5, misi_title=$6, misi_text=$7, 
			misi_text2=$8, nilai_title=$9, nilai_title1=$10, nilai_text1=$11, nilai_image1=$12, nilai_title2=$13, nilai_text2=$14, nilai_image2=$15, nilai_title3=$16, nilai_text3=$17, 
			nilai_image3=$18, nilai_title4=$19, nilai_text4=$20, nilai_image4=$21, struktur_title=$22, struktur_image=$23, updated_at=NOW() WHERE id=$24`
	if _, err := tx.ExecContext(ctx, sql, front.LatarTitle, front.LatarText, front.VisiMisiTitle, front.VisiTitle, front.VisiText, front.MisiTitle, front.MisiText, front.MisiText2, front.NilaiTitle, front.NilaiTitle1, front.NilaiText1, front.NilaiImage1, front.NilaiTitle2, front.NilaiText2, front.NilaiImage2, front.NilaiTitle3, front.NilaiText3, front.NilaiImage3, front.NilaiTitle4, front.NilaiText4, front.NilaiImage4, front.StrukturTitle, front.StrukturImage, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *AboutRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, front *entity.About) error {
	sql := `DELETE FROM abouts WHERE id=$1`
	if _, err := tx.ExecContext(ctx, sql, front.Id); err != nil {
		return err
	}
	return nil
}

func (r *AboutRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.About, error) {
	var front entity.About
	sql := `SELECT id, latar_title, latar_text, visi_misi_title, visi_title, visi_text, misi_title, misi_text, misi_text2, nilai_title, nilai_title1, nilai_text1, nilai_image1, nilai_title2, nilai_text2, nilai_image2, nilai_title3, nilai_text3, nilai_image3, nilai_title4, nilai_text4, nilai_image4, struktur_title, struktur_image, created_at, updated_at FROM abouts WHERE id=$1`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.LatarTitle, &front.LatarText, &front.VisiMisiTitle, &front.VisiTitle, &front.VisiText, &front.MisiTitle, &front.MisiText, &front.MisiText2, &front.NilaiTitle, &front.NilaiTitle1, &front.NilaiText1, &front.NilaiImage1, &front.NilaiTitle2, &front.NilaiText2, &front.NilaiImage2, &front.NilaiTitle3, &front.NilaiText3, &front.NilaiImage3, &front.NilaiTitle4, &front.NilaiText4, &front.NilaiImage4, &front.StrukturTitle, &front.StrukturImage, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *AboutRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.About, error) {
	var fronts []entity.About
	sql := `SELECT id, latar_title, latar_text, visi_misi_title, visi_title, visi_text, misi_title, misi_text, misi_text2, nilai_title, nilai_title1, nilai_text1, nilai_image1, nilai_title2, nilai_text2, nilai_image2, nilai_title3, nilai_text3, nilai_image3, nilai_title4, nilai_text4, nilai_image4, struktur_title, struktur_image, created_at, updated_at FROM abouts ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return fronts, err
	}
	defer rows.Close()

	for rows.Next() {
		var front entity.About
		if err := rows.Scan(&front.Id, &front.LatarTitle, &front.LatarText, &front.LatarTitle, &front.LatarText, &front.VisiMisiTitle, &front.VisiTitle, &front.VisiText, &front.MisiTitle, &front.MisiText, &front.MisiText2, &front.NilaiTitle, &front.NilaiTitle1, &front.NilaiText1, &front.NilaiImage1, &front.NilaiTitle2, &front.NilaiText2, &front.NilaiImage2, &front.NilaiTitle3, &front.NilaiText3, &front.NilaiImage3, &front.NilaiTitle4, &front.NilaiText4, &front.NilaiImage4, &front.StrukturTitle, &front.StrukturImage); err != nil {
			return fronts, err
		}
		fronts = append(fronts, front)
	}

	return fronts, nil
}

func (r *AboutRepositoryImpl) FindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM abouts`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
