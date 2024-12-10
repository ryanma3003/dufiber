package dto

type BlogCreate struct {
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	Content        string `json:"content"`
	Image          string `json:"image"`
	Author         string `json:"author"`
	UserId         int    `json:"user_id"`
	BlogCategoryId int    `json:"blog_category_id"`
}

type BlogUpdate struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	Content        string `json:"content"`
	Image          string `json:"image"`
	Author         string `json:"author"`
	UserId         int    `json:"user_id"`
	BlogCategoryId int    `json:"blog_category_id"`
}

type BlogResponse struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	Content        string `json:"content"`
	Image          string `json:"image"`
	Author         string `json:"author"`
	UserId         int    `json:"user_id"`
	BlogCategoryId int    `json:"blog_category_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type BlogCategoryCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BlogCategoryUpdate struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BlogCategoryResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
