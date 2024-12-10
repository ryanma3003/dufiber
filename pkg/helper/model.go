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
