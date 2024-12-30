package helper

import (
	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
)

// User

func ToUserResponse(user entity.User) dto.UserResponse {
	return dto.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []entity.User) []dto.UserResponse {
	var userRes []dto.UserResponse

	if users == nil {
		return []dto.UserResponse{}
	}

	for _, user := range users {
		userRes = append(userRes, ToUserResponse(user))
	}

	return userRes
}

// Blog

func ToBlogResponse(blog entity.Blog) dto.BlogResponse {
	return dto.BlogResponse{
		Id:             blog.Id,
		Title:          blog.Title,
		Slug:           blog.Slug,
		Content:        blog.Content,
		Image:          blog.Image,
		Author:         blog.Author,
		UserId:         blog.UserId,
		BlogCategoryId: blog.BlogCategoryId,
		CreatedAt:      blog.CreatedAt,
		UpdatedAt:      blog.UpdatedAt,
	}
}

func ToBlogResponses(blogs []entity.Blog) []dto.BlogResponse {
	var blogRes []dto.BlogResponse

	if blogs == nil {
		return []dto.BlogResponse{}
	}

	for _, blog := range blogs {
		blogRes = append(blogRes, ToBlogResponse(blog))
	}

	return blogRes
}

// Blog category

func ToBlogCategoryResponse(blog entity.BlogCategory) dto.BlogCategoryResponse {
	return dto.BlogCategoryResponse{
		Id:          blog.Id,
		Title:       blog.Title,
		Description: blog.Description,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}
}

func ToBlogCategoryResponses(blogs []entity.BlogCategory) []dto.BlogCategoryResponse {
	var blogRes []dto.BlogCategoryResponse

	if blogs == nil {
		return []dto.BlogCategoryResponse{}
	}

	for _, blog := range blogs {
		blogRes = append(blogRes, ToBlogCategoryResponse(blog))
	}

	return blogRes
}

// Donation
func ToDonationResponse(donation entity.Donation) dto.DonationResponse {
	return dto.DonationResponse{
		Id:             donation.Id,
		Name:           donation.Name,
		Email:          donation.Email,
		Phone:          donation.Phone,
		Amount:         donation.Amount,
		Status:         donation.Status,
		Reference:      donation.Reference,
		SnapToken:      donation.SnapToken,
		DonationListId: donation.DonationListId,
		CharityListId:  donation.CharityListId,
		UserId:         donation.UserId,
		OrderId:        donation.OrderId,
		CreatedAt:      donation.CreatedAt,
		UpdatedAt:      donation.UpdatedAt,
	}
}

func ToDonationResponses(donations []entity.Donation) []dto.DonationResponse {
	var donationRes []dto.DonationResponse

	if donations == nil {
		return []dto.DonationResponse{}
	}

	for _, donation := range donations {
		donationRes = append(donationRes, ToDonationResponse(donation))
	}

	return donationRes
}

// Donation Category
func ToDonationCategoryResponse(donation entity.DonationCategory) dto.DonationCategoryResponse {
	return dto.DonationCategoryResponse{
		Id:          donation.Id,
		Title:       donation.Title,
		Description: donation.Description,
		CreatedAt:   donation.CreatedAt,
		UpdatedAt:   donation.UpdatedAt,
	}
}

func ToDonationCategoryResponses(donations []entity.DonationCategory) []dto.DonationCategoryResponse {
	var donationRes []dto.DonationCategoryResponse

	if donations == nil {
		return []dto.DonationCategoryResponse{}
	}

	for _, donation := range donations {
		donationRes = append(donationRes, ToDonationCategoryResponse(donation))
	}

	return donationRes
}

// Donation List
func ToDonationListResponse(donation entity.DonationList) dto.DonationListResponse {
	return dto.DonationListResponse{
		Id:                 donation.Id,
		Title:              donation.Title,
		Description:        donation.Description,
		Code:               donation.Code,
		DonationCategoryId: donation.DonationCategoryId,
		CreatedAt:          donation.CreatedAt,
		UpdatedAt:          donation.UpdatedAt,
	}
}

func ToDonationListResponses(donations []entity.DonationList) []dto.DonationListResponse {
	var donationRes []dto.DonationListResponse

	if donations == nil {
		return []dto.DonationListResponse{}
	}

	for _, donation := range donations {
		donationRes = append(donationRes, ToDonationListResponse(donation))
	}

	return donationRes
}

// Harga Zakat
func ToHargaZakatResponse(donation entity.HargaZakat) dto.HargaZakatResponse {
	return dto.HargaZakatResponse{
		Id:             donation.Id,
		Title:          donation.Title,
		Price:          donation.Price,
		DonationListId: donation.DonationListId,
		CreatedAt:      donation.CreatedAt,
		UpdatedAt:      donation.UpdatedAt,
	}
}

func ToHargaZakatResponses(donations []entity.HargaZakat) []dto.HargaZakatResponse {
	var donationRes []dto.HargaZakatResponse

	if donations == nil {
		return []dto.HargaZakatResponse{}
	}

	for _, donation := range donations {
		donationRes = append(donationRes, ToHargaZakatResponse(donation))
	}

	return donationRes
}

// Galeri Tag
func ToGaleriTagResponse(data entity.GaleriTag) dto.GaleriTagResponse {
	return dto.GaleriTagResponse{
		Id:        data.Id,
		Title:     data.Title,
		Slug:      data.Slug,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToGaleriTagResponses(datas []entity.GaleriTag) []dto.GaleriTagResponse {
	var dataRes []dto.GaleriTagResponse

	if datas == nil {
		return []dto.GaleriTagResponse{}
	}

	for _, data := range datas {
		dataRes = append(dataRes, ToGaleriTagResponse(data))
	}

	return dataRes
}

// Galeri
func ToGaleriResponse(data entity.Galeri) dto.GaleriResponse {
	return dto.GaleriResponse{
		Id:          data.Id,
		Title:       data.Title,
		Slug:        data.Slug,
		Image:       data.Image,
		GaleryTagId: data.GaleryTagId,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ToGaleriResponses(datas []entity.Galeri) []dto.GaleriResponse {
	var dataRes []dto.GaleriResponse

	if datas == nil {
		return []dto.GaleriResponse{}
	}

	for _, data := range datas {
		dataRes = append(dataRes, ToGaleriResponse(data))
	}

	return dataRes
}

// Homepage
func ToHomepageResponse(data entity.Homepage) dto.HomepageResponse {
	return dto.HomepageResponse{
		Id:              data.Id,
		MainImage:       data.MainImage,
		MainTitle:       data.MainTitle,
		MainText:        data.MainText,
		KalkulatorTitle: data.KalkulatorTitle,
		KalkulatorText:  data.KalkulatorText,
		PersText:        data.PersText,
		PublikasiText:   data.PublikasiText,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}
}

func ToHomepageResponses(datas []entity.Homepage) []dto.HomepageResponse {
	var dataRes []dto.HomepageResponse

	if datas == nil {
		return []dto.HomepageResponse{}
	}

	for _, data := range datas {
		dataRes = append(dataRes, ToHomepageResponse(data))
	}

	return dataRes
}

// About
func ToAboutResponse(data entity.About) dto.AboutResponse {
	return dto.AboutResponse{
		Id:            data.Id,
		LatarTitle:    data.LatarTitle,
		LatarText:     data.LatarText,
		VisiMisiTitle: data.VisiMisiTitle,
		VisiTitle:     data.VisiTitle,
		VisiText:      data.VisiText,
		MisiTitle:     data.MisiTitle,
		MisiText:      data.MisiText,
		MisiText2:     data.MisiText2,
		NilaiTitle:    data.NilaiTitle,
		NilaiTitle1:   data.NilaiTitle1,
		NilaiText1:    data.NilaiText1,
		NilaiImage1:   data.NilaiImage1,
		NilaiTitle2:   data.NilaiTitle2,
		NilaiText2:    data.NilaiText2,
		NilaiImage2:   data.NilaiImage2,
		NilaiTitle3:   data.NilaiTitle3,
		NilaiText3:    data.NilaiText3,
		NilaiImage3:   data.NilaiImage3,
		NilaiTitle4:   data.NilaiTitle4,
		NilaiText4:    data.NilaiText4,
		NilaiImage4:   data.NilaiImage4,
		StrukturTitle: data.StrukturTitle,
		StrukturImage: data.StrukturImage,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
}

func ToAboutResponses(datas []entity.About) []dto.AboutResponse {
	var dataRes []dto.AboutResponse

	if datas == nil {
		return []dto.AboutResponse{}
	}

	for _, data := range datas {
		dataRes = append(dataRes, ToAboutResponse(data))
	}

	return dataRes
}

// Contact
func ToContactResponse(data entity.Contact) dto.ContactResponse {
	return dto.ContactResponse{
		Id:           data.Id,
		MainTag:      data.MainTag,
		MainText:     data.MainText,
		AddressTitle: data.AddressTitle,
		Address:      data.Address,
		PhoneTitle:   data.PhoneTitle,
		Phone:        data.Phone,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToContactResponses(datas []entity.Contact) []dto.ContactResponse {
	var dataRes []dto.ContactResponse

	if datas == nil {
		return []dto.ContactResponse{}
	}

	for _, data := range datas {
		dataRes = append(dataRes, ToContactResponse(data))
	}

	return dataRes
}

// FAQ
func ToFaqResponse(data entity.Faq) dto.FaqResponse {
	return dto.FaqResponse{
		Id:         data.Id,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		Pertanyaan: data.Pertanyaan,
		Jawaban:    data.Jawaban,
	}
}

func ToFaqResponses(datas []entity.Faq) []dto.FaqResponse {
	var dataRes []dto.FaqResponse

	if datas == nil {
		return []dto.FaqResponse{}
	}

	for _, data := range datas {
		dataRes = append(dataRes, ToFaqResponse(data))
	}

	return dataRes
}

// Term
func ToTermResponse(data entity.Term) dto.TermResponse {
	return dto.TermResponse{
		Id:        data.Id,
		Title:     data.Title,
		Text:      data.Text,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToTermResponses(datas []entity.Term) []dto.TermResponse {
	var dataRes []dto.TermResponse

	if datas == nil {
		return []dto.TermResponse{}
	}

	for _, data := range datas {
		dataRes = append(dataRes, ToTermResponse(data))
	}

	return dataRes
}

// Privacy
func ToPrivacyResponse(data entity.Privacy) dto.PrivacyResponse {
	return dto.PrivacyResponse{
		Id:        data.Id,
		Title:     data.Title,
		Text:      data.Text,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToPrivacyResponses(data []entity.Privacy) []dto.PrivacyResponse {
	var dataRes []dto.PrivacyResponse

	if data == nil {
		return []dto.PrivacyResponse{}
	}

	for _, data := range data {
		dataRes = append(dataRes, ToPrivacyResponse(data))
	}

	return dataRes
}

// Ikrar
func ToIkrarResponse(data entity.Ikrar) dto.IkrarResponse {
	return dto.IkrarResponse{
		Id:             data.Id,
		Nama:           data.Nama,
		Email:          data.Email,
		Telepon:        data.Telepon,
		Tanggal:        data.Tanggal,
		NamaHari:       data.NamaHari,
		JumlahDonasi:   data.JumlahDonasi,
		JumlahPohon:    data.JumlahPohon,
		HargaSatuPohon: data.HargaSatuPohon,
		NamaPohon:      data.NamaPohon,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}
}

func ToIkrarResponses(data []entity.Ikrar) []dto.IkrarResponse {
	var dataRes []dto.IkrarResponse

	if data == nil {
		return []dto.IkrarResponse{}
	}

	for _, data := range data {
		dataRes = append(dataRes, ToIkrarResponse(data))
	}

	return dataRes
}
