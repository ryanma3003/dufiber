package entity

type GaleriTag struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Slug      string `json:"slug"`
}

type Galeri struct {
	Id          int    `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	GaleriTagId int    `json:"galeri_tag_id"`
}
