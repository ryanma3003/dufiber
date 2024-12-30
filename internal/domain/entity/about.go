package entity

type About struct {
	Id            int     `json:"id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	LatarTitle    string  `json:"latar_title"`
	LatarText     string  `json:"latar_text"`
	VisiMisiTitle string  `json:"visi_misi_title"`
	VisiTitle     string  `json:"visi_title"`
	VisiText      string  `json:"visi_text"`
	MisiTitle     string  `json:"misi_title"`
	MisiText      string  `json:"misi_text"`
	MisiText2     string  `json:"misi_text2"`
	NilaiTitle    *string `json:"nilai_title"`
	NilaiTitle1   *string `json:"nilai_title1"`
	NilaiText1    *string `json:"nilai_text1"`
	NilaiImage1   string  `json:"nilai_image1"`
	NilaiTitle2   *string `json:"nilai_title2"`
	NilaiText2    *string `json:"nilai_text2"`
	NilaiImage2   string  `json:"nilai_image2"`
	NilaiTitle3   *string `json:"nilai_title3"`
	NilaiText3    *string `json:"nilai_text3"`
	NilaiImage3   string  `json:"nilai_image3"`
	NilaiTitle4   *string `json:"nilai_title4"`
	NilaiText4    *string `json:"nilai_text4"`
	NilaiImage4   string  `json:"nilai_image4"`
	StrukturTitle string  `json:"struktur_title"`
	StrukturImage *string `json:"struktur_image"`
}
