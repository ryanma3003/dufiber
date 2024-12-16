package dto

// all landing pages are here
// Homepage Page

type HomepageCreate struct {
	MainImage       string `json:"main_image"`
	MainText        string `json:"main_text"`
	MainTitle       string `json:"main_title"`
	KalkulatorTitle string `json:"kalkulator_title"`
	KalkulatorText  string `json:"kalkulator_text"`
	PersText        string `json:"pers_text"`
	PublikasiText   string `json:"publikasi_text"`
}

type HomepageUpdate struct {
	Id              int    `json:"id"`
	MainImage       string `json:"main_image"`
	MainText        string `json:"main_text"`
	MainTitle       string `json:"main_title"`
	KalkulatorTitle string `json:"kalkulator_title"`
	KalkulatorText  string `json:"kalkulator_text"`
	PersText        string `json:"pers_text"`
	PublikasiText   string `json:"publikasi_text"`
}

type HomepageResponse struct {
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

// About Page

type AboutCreate struct {
	LatarTitle    string `json:"latar_title"`
	LatarText     string `json:"latar_text"`
	VisiMisiTitle string `json:"visi_misi_title"`
	VisiTitle     string `json:"visi_title"`
	VisiText      string `json:"visi_text"`
	MisiTitle     string `json:"misi_title"`
	MisiText      string `json:"misi_text"`
	MisiText2     string `json:"misi_text2"`
	NilaiTitle    string `json:"nilai_title"`
	NilaiTitle1   string `json:"nilai_title1"`
	NilaiText1    string `json:"nilai_text1"`
	NilaiImage1   string `json:"nilai_image1"`
	NilaiTitle2   string `json:"nilai_title2"`
	NilaiText2    string `json:"nilai_text2"`
	NilaiImage2   string `json:"nilai_image2"`
	NilaiTitle3   string `json:"nilai_title3"`
	NilaiText3    string `json:"nilai_text3"`
	NilaiImage3   string `json:"nilai_image3"`
	NilaiTitle4   string `json:"nilai_title4"`
	NilaiText4    string `json:"nilai_text4"`
	NilaiImage4   string `json:"nilai_image4"`
	StrukturTitle string `json:"struktur_title"`
	StrukturImage string `json:"struktur_image"`
}

type AboutUpdate struct {
	Id            int    `json:"id"`
	LatarTitle    string `json:"latar_title"`
	LatarText     string `json:"latar_text"`
	VisiMisiTitle string `json:"visi_misi_title"`
	VisiTitle     string `json:"visi_title"`
	VisiText      string `json:"visi_text"`
	MisiTitle     string `json:"misi_title"`
	MisiText      string `json:"misi_text"`
	MisiText2     string `json:"misi_text2"`
	NilaiTitle    string `json:"nilai_title"`
	NilaiTitle1   string `json:"nilai_title1"`
	NilaiText1    string `json:"nilai_text1"`
	NilaiImage1   string `json:"nilai_image1"`
	NilaiTitle2   string `json:"nilai_title2"`
	NilaiText2    string `json:"nilai_text2"`
	NilaiImage2   string `json:"nilai_image2"`
	NilaiTitle3   string `json:"nilai_title3"`
	NilaiText3    string `json:"nilai_text3"`
	NilaiImage3   string `json:"nilai_image3"`
	NilaiTitle4   string `json:"nilai_title4"`
	NilaiText4    string `json:"nilai_text4"`
	NilaiImage4   string `json:"nilai_image4"`
	StrukturTitle string `json:"struktur_title"`
	StrukturImage string `json:"struktur_image"`
}

type AboutResponse struct {
	Id            int    `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	LatarTitle    string `json:"latar_title"`
	LatarText     string `json:"latar_text"`
	VisiMisiTitle string `json:"visi_misi_title"`
	VisiTitle     string `json:"visi_title"`
	VisiText      string `json:"visi_text"`
	MisiTitle     string `json:"misi_title"`
	MisiText      string `json:"misi_text"`
	MisiText2     string `json:"misi_text2"`
	NilaiTitle    string `json:"nilai_title"`
	NilaiTitle1   string `json:"nilai_title1"`
	NilaiText1    string `json:"nilai_text1"`
	NilaiImage1   string `json:"nilai_image1"`
	NilaiTitle2   string `json:"nilai_title2"`
	NilaiText2    string `json:"nilai_text2"`
	NilaiImage2   string `json:"nilai_image2"`
	NilaiTitle3   string `json:"nilai_title3"`
	NilaiText3    string `json:"nilai_text3"`
	NilaiImage3   string `json:"nilai_image3"`
	NilaiTitle4   string `json:"nilai_title4"`
	NilaiText4    string `json:"nilai_text4"`
	NilaiImage4   string `json:"nilai_image4"`
	StrukturTitle string `json:"struktur_title"`
	StrukturImage string `json:"struktur_image"`
}

// Contact Page

type ContactCreate struct {
	MainTag      string `json:"main_tag"`
	MainText     string `json:"main_text"`
	AddressTitle string `json:"address_title"`
	Address      string `json:"address"`
	PhoneTitle   string `json:"phone_title"`
	Phone        string `json:"phone"`
}

type ContactUpdate struct {
	Id           int    `json:"id"`
	MainTag      string `json:"main_tag"`
	MainText     string `json:"main_text"`
	AddressTitle string `json:"address_title"`
	Address      string `json:"address"`
	PhoneTitle   string `json:"phone_title"`
	Phone        string `json:"phone"`
}

type ContactResponse struct {
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

// FAQ Page

type FaqCreate struct {
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
}

type FaqUpdate struct {
	Id         int    `json:"id"`
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
}

type FaqResponse struct {
	Id         int    `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
}

// Term and Privacy Page

type TermCreate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type TermUpdate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type TermResponse struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Text      string `json:"text"`
}

type PrivacyCreate struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PrivacyUpdate struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type PrivacyResponse struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Text      string `json:"text"`
}
