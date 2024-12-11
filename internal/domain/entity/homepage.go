package entity

type Homepage struct {
	Id              int    `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	MainImage       string `json:"main_image"`
	MainText        string `json:"main_text"`
	MainTitle       string `json:"main_title"`
	KalkulatorTitle string `json:"kalkulator_title"`
	KalkulatorText  string `json:"kalkulator_text"`
	PersText        string `json:"pers_text"`
	PublikasiText   string `json:"publikasi_text"`
}
