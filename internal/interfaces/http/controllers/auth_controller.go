package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService}
}

func (h *AuthController) Login(c *fiber.Ctx) error {
	loginInput := new(dto.LoginInput)
	if err := c.BodyParser(loginInput); err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := helper.ValidateStruct(loginInput); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Username":
				return helper.RespondError(c, fiber.StatusBadRequest, "Invalid username")
			case "Password":
				return helper.RespondError(c, fiber.StatusBadRequest, "Invalid password")
			default:
				return helper.RespondError(c, fiber.StatusBadRequest, err.Error())
			}
		}

	}

	token, err := h.authService.LoginUser(c.Context(), loginInput)
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusBadRequest, err.Error())
	}

	return helper.RespondWithData(c, fiber.StatusOK, "Login success", fiber.Map{
		"token":        token.Token,
		"token_type":   "Bearer",
		"expired_time": "8h",
	})
}
