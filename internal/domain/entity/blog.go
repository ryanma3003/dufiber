package entity

type BlogCategory struct {
	Id          int    `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Blog struct {
	Id             int    `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	Content        string `json:"content"`
	Image          string `json:"image"`
	Author         string `json:"author"`
	UserId         int    `json:"user_id"`
	BlogCategoryId int    `json:"blog_category_id"`
}
