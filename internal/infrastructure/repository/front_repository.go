package repository

import (
	"context"
	"database/sql"

	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/domain/repository"
)

type FrontRepositoryImpl struct{}

func NewFrontRepository() repository.FrontRepository {
	return &FrontRepositoryImpl{}
}

func (r *FrontRepositoryImpl) HomepageFindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Homepage, error) {
	var front entity.Homepage
	sql := `SELECT id, main_image, main_title, main_text, kalkulator_title, kalkulator_text, pers_text, publikasi_text, created_at, updated_at FROM homepages WHERE id=?`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.MainImage, &front.MainTitle, &front.MainText, &front.KalkulatorTitle, &front.KalkulatorText, &front.PersText, &front.PublikasiText, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *FrontRepositoryImpl) AboutFindByID(ctx context.Context, tx *sql.Tx, id int) (entity.About, error) {
	var front entity.About
	sql := `SELECT id, latar_title, latar_text, visi_misi_title, visi_title, visi_text, misi_title, misi_text, misi_text2, nilai_title, nilai_title1, nilai_text1, nilai_image1, nilai_title2, nilai_text2, nilai_image2, nilai_title3, nilai_text3, nilai_image3, nilai_title4, nilai_text4, nilai_image4, struktur_title, struktur_image, created_at, updated_at FROM abouts WHERE id=?`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.LatarTitle, &front.LatarText, &front.VisiMisiTitle, &front.VisiTitle, &front.VisiText, &front.MisiTitle, &front.MisiText, &front.MisiText2, &front.NilaiTitle, &front.NilaiTitle1, &front.NilaiText1, &front.NilaiImage1, &front.NilaiTitle2, &front.NilaiText2, &front.NilaiImage2, &front.NilaiTitle3, &front.NilaiText3, &front.NilaiImage3, &front.NilaiTitle4, &front.NilaiText4, &front.NilaiImage4, &front.StrukturTitle, &front.StrukturImage, &front.CreatedAt, &front.UpdatedAt); err != nil {
		return front, err
	}
	return front, nil
}

func (r *FrontRepositoryImpl) GaleriFindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Galeri, error) {
	var galeris []entity.Galeri
	sql := `SELECT id, title, slug, image, galeri_tag_id, created_at, updated_at FROM galeris ORDER BY created_at DESC LIMIT ? OFFSET ?`
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		return galeris, err
	}
	defer rows.Close()

	for rows.Next() {
		var galeri entity.Galeri
		if err := rows.Scan(&galeri.Id, &galeri.Title, &galeri.Slug, galeri.Image, galeri.GaleriTagId, &galeri.CreatedAt, &galeri.UpdatedAt); err != nil {
			return galeris, err
		}
		galeris = append(galeris, galeri)
	}

	return galeris, nil
}

func (r *FrontRepositoryImpl) GaleriFindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM galeris`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}

func (r *FrontRepositoryImpl) FaqFindAll(ctx context.Context, tx *sql.Tx) ([]entity.Faq, error) {
	var fronts []entity.Faq
	sql := `SELECT id, created_at, updated_at, pertanyaan, jawaban FROM faq ORDER BY created_at DESC`
	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		return fronts, err
	}
	defer rows.Close()

	for rows.Next() {
		var front entity.Faq
		if err := rows.Scan(&front.Id, &front.CreatedAt, &front.UpdatedAt, &front.Pertanyaan, &front.Jawaban); err != nil {
			return fronts, err
		}
		fronts = append(fronts, front)
	}

	return fronts, nil
}

func (r *FrontRepositoryImpl) TermFindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Term, error) {
	var front entity.Term
	sql := `SELECT id, created_at, updated_at, title, text FROM terms WHERE id=?`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.CreatedAt, &front.UpdatedAt, &front.Title, &front.Text); err != nil {
		return front, err
	}
	return front, nil
}

func (r *FrontRepositoryImpl) PrivacyFindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Privacy, error) {
	var front entity.Privacy
	sql := `SELECT id, created_at, updated_at, title, text FROM privacies WHERE id=?`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.CreatedAt, &front.UpdatedAt, &front.Title, &front.Text); err != nil {
		return front, err
	}
	return front, nil
}

func (r *FrontRepositoryImpl) ContactFindByID(ctx context.Context, tx *sql.Tx, id int) (entity.Contact, error) {
	var front entity.Contact
	sql := `SELECT id, created_at, updated_at, main_tag, main_text, address_title, address, phone_title, phone FROM contacts WHERE id=?`
	if err := tx.QueryRowContext(ctx, sql, id).Scan(&front.Id, &front.CreatedAt, &front.UpdatedAt, &front.MainTag, &front.MainText, &front.AddressTitle, &front.Address, &front.PhoneTitle, &front.Phone); err != nil {
		return front, err
	}
	return front, nil
}

func (r *FrontRepositoryImpl) BlogFindBySlug(ctx context.Context, tx *sql.Tx, slug string) (entity.Blog, error) {
	var blog entity.Blog
	sql := `SELECT id, created_at, updated_at, title, slug, content, image, author, user_id, blog_category_id FROM blog WHERE slug=?`
	if err := tx.QueryRowContext(ctx, sql, slug).Scan(&blog.Id, &blog.CreatedAt, &blog.UpdatedAt, &blog.Title, &blog.Slug, &blog.Content, &blog.Image, &blog.Author, &blog.UserId, &blog.BlogCategoryId); err != nil {
		return blog, err
	}
	return blog, nil
}

func (r *FrontRepositoryImpl) BlogFindAllWithPagination(ctx context.Context, tx *sql.Tx, limit, offset int) ([]entity.Blog, error) {
	var blogs []entity.Blog
	sql := `SELECT id, title, image, content, slug, author, blog_category_id, user_id, created_at, updated_at FROM blog ORDER BY created_at DESC LIMIT ? OFFSET $2`
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

func (r *FrontRepositoryImpl) BlogFindTotal(ctx context.Context, tx *sql.Tx) (int, error) {
	var total int
	sql := `SELECT COUNT(*) FROM blog`
	if err := tx.QueryRowContext(ctx, sql).Scan(&total); err != nil {
		return total, err
	}
	return total, nil
}
