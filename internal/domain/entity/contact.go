package entity

type Contact struct {
	Id           int    `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	MainTag      string `json:"main_tag"`
	MainText     string `json:"main_text"`
	AddressTitle string `json:"address_title"`
	Address      string `json:"address"`
	PhoneTitle   string `json:"phone_title"`
	Phone        string `json:"phone"`
}
