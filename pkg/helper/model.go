package helper

import (
	"github.com/ryanma3003/dufiber/internal/domain/entity"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
)

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

func ToBlogResponse(blog entity.Blog) dto.BlogResponse {
	return dto.BlogResponse{
		Id:             blog.Id,
		Title:          blog.Title,
		Slug:           blog.Slug,
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
