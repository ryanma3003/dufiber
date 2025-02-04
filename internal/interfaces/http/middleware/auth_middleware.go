package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ryanma3003/dufiber/internal/infrastructure/database"
	"github.com/ryanma3003/dufiber/internal/infrastructure/repository"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

func WebAuth(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c) // Get session ( creates one if not exist )
		if err != nil {
			return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, err.Error())
		}

		// Get user ID from session
		userID := sess.Get("user_id")
		if userID == nil {
			return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, "cant Get user ID from session")
		}

		// check the user
		user_service := service.NewUserService(repository.NewUserRepository(), database.DB)
		user, err := user_service.FindById(c.Context(), userID.(int))
		if err != nil {
			return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, err.Error())
		}

		userSession := dto.UserSession{
			Id:       user.Id,
			Username: user.Username,
			Role:     user.Role,
		}

		c.Locals("user", userSession)

		return c.Next()
	}
}

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
