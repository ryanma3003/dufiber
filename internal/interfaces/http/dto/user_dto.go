package dto

type UserCreate struct {
	Username string `json:"username" validate:"required,min=5,max=50,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50,containsany=!@#$%^&*()_+,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
	Role     int    `json:"role"`
}

type UserUpdate struct {
	Id       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required,min=5,max=50,alphanum"`
	Role     int    `json:"role"`
	Email    string `json:"email" validate:"required,email"`
}

type UserChangePassword struct {
	Id          int    `json:"id"`
	OldPassword string `json:"old_password" validate:"required"`
	Password    string `json:"password" validate:"required,min=6,max=50,containsany=!@#$%^&*()_+,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserSession struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}
