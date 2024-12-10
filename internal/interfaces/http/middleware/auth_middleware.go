package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ryanma3003/dufiber/internal/infrastructure/database"
	"github.com/ryanma3003/dufiber/internal/infrastructure/repository"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

func IsAuth(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	headerSplit := strings.Split(header, "Bearer ")
	if len(headerSplit) != 2 {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	token := headerSplit[1]
	if token == "" {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	// decode token
	decode_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SERCRET")), nil
	})
	if err != nil {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	id := decode_token.Claims.(jwt.MapClaims)["id"].(float64)

	user_service := service.NewUserService(repository.NewUserRepository(), database.DB)
	user, err := user_service.FindById(c.Context(), int(id))
	if err != nil {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userSession := dto.UserSession{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
	}

	c.Locals("user", userSession)

	return c.Next()
}
