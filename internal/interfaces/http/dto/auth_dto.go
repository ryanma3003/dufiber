package dto

type LoginInput struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Id        int    `json:"id"`
	Token     string `json:"token"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
