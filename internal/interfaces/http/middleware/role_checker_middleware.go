package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

func IsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(dto.UserSession)

	if user.Role != 1 {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized: not an admin")
	}

	return c.Next()
}

func IsSuperAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(dto.UserSession)

	if user.Role != 3 {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized: not a super admin")
	}

	return c.Next()
}

func IsSelf(c *fiber.Ctx) error {
	user := c.Locals("user").(dto.UserSession)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	if user.Id != id {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized: id not valid")
	}

	return c.Next()
}

func IsSuperAdminOrIsSelf(c *fiber.Ctx) error {
	user := c.Locals("user").(dto.UserSession)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	if user.Role != 3 && user.Id != id {
		return helper.RespondError(c, fiber.StatusUnauthorized, "Unauthorized: not a super admin or not the user")
	}

	return c.Next()
}
