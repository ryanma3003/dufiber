package dto

// Galeri DTO

type GaleriTagCreate struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type GaleriTagUpdate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type GaleriTagResponse struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Slug      string `json:"slug"`
}

type GaleriCreate struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	GaleryTagId int    `json:"galeri_tag_id"`
}

type GaleriUpdate struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	GaleryTagId int    `json:"galeri_tag_id"`
}

type GaleriResponse struct {
	Id          int    `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	GaleryTagId int    `json:"galeri_tag_id"`
}
