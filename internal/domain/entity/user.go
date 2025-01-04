package entity

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`
	DeletedAt string `json:"deleted_at"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
