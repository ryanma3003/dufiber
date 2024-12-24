package dto

type IkrarCreate struct {
	Nama           string `json:"nama" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Telepon        string `json:"telepon" validate:"required"`
	Tanggal        string `json:"tanggal" validate:"required"`
	NamaHari       string `json:"nama_hari" validate:"required"`
	JumlahDonasi   int    `json:"jumlah_donasi" validate:"required"`
	JumlahPohon    int    `json:"jumlah_pohon" validate:"required"`
	HargaSatuPohon int    `json:"harga_satu_pohon" validate:"required"`
	NamaPohon      string `json:"nama_pohon" validate:"required"`
}

type IkrarUpdate struct {
	Id             int    `json:"id" validate:"required"`
	Nama           string `json:"nama" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Telepon        string `json:"telepon" validate:"required"`
	Tanggal        string `json:"tanggal" validate:"required"`
	NamaHari       string `json:"nama_hari" validate:"required"`
	JumlahDonasi   int    `json:"jumlah_donasi" validate:"required"`
	JumlahPohon    int    `json:"jumlah_pohon" validate:"required"`
	HargaSatuPohon int    `json:"harga_satu_pohon" validate:"required"`
	NamaPohon      string `json:"nama_pohon" validate:"required"`
}

type IkrarResponse struct {
	Id             int    `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Nama           string `json:"nama"`
	Email          string `json:"email"`
	Telepon        string `json:"telepon"`
	Tanggal        string `json:"tanggal"`
	NamaHari       string `json:"nama_hari"`
	JumlahDonasi   int    `json:"jumlah_donasi"`
	JumlahPohon    int    `json:"jumlah_pohon"`
	HargaSatuPohon int    `json:"harga_satu_pohon"`
	NamaPohon      string `json:"nama_pohon"`
}
