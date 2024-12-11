package entity

type Faq struct {
	Id         int    `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
}
