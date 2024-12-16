package dto

type IkrarCreate struct {
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

type IkrarUpdate struct {
	Id             int    `json:"id"`
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
