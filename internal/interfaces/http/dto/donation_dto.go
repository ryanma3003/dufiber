package dto

// All donation related DTOs are here

type DonationCategoryCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DonationCategoryUpdate struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DonationCategoryResponse struct {
	Id          int    `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DonationListCreate struct {
	Title              string `json:"title"`
	Description        string `json:"description"`
	Code               int    `json:"code"`
	DonationCategoryId int    `json:"donation_category_id"`
}

type DonationListUpdate struct {
	Id                 int    `json:"id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Code               int    `json:"code"`
	DonationCategoryId int    `json:"donation_category_id"`
}

type DonationListResponse struct {
	Id                 int    `json:"id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Code               int    `json:"code"`
	DonationCategoryId int    `json:"donation_category_id"`
}

type DonationCreate struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Amount         int    `json:"amount"`
	Status         string `json:"status"`
	Reference      string `json:"reference"`
	SnapToken      string `json:"snap_token"`
	DonationListId int    `json:"donation_list_id"`
	CharityListId  int    `json:"charity_list_id"`
	UserId         int    `json:"user_id"`
	OrderId        string `json:"order_id"`
}

type DonationUpdate struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Amount         int    `json:"amount"`
	Status         string `json:"status"`
	Reference      string `json:"reference"`
	SnapToken      string `json:"snap_token"`
	DonationListId int    `json:"donation_list_id"`
	CharityListId  int    `json:"charity_list_id"`
	UserId         int    `json:"user_id"`
	OrderId        string `json:"order_id"`
}

type DonationResponse struct {
	Id             int    `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Amount         int    `json:"amount"`
	Status         string `json:"status"`
	Reference      string `json:"reference"`
	SnapToken      string `json:"snap_token"`
	DonationListId int    `json:"donation_list_id"`
	CharityListId  int    `json:"charity_list_id"`
	UserId         int    `json:"user_id"`
	OrderId        string `json:"order_id"`
}

type HargaZakatCreate struct {
	DonationListId int    `json:"donation_list_id"`
	Title          string `json:"title"`
	Price          int    `json:"price"`
}

type HargaZakatUpdate struct {
	Id             int    `json:"id"`
	DonationListId int    `json:"donation_list_id"`
	Title          string `json:"title"`
	Price          int    `json:"price"`
}

type HargaZakatResponse struct {
	Id             int    `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DonationListId int    `json:"donation_list_id"`
	Title          string `json:"title"`
	Price          int    `json:"price"`
}
