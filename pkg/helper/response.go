package helper

import "github.com/gofiber/fiber/v2"

func RespondMessage(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   false,
		"message": message,
	})
}

func RespondWithData(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   false,
		"message": message,
		"data":    data,
	})
}

func RespondWithPagination(c *fiber.Ctx, code int, message string, total int, page int, perPage int, dataName string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"error":   false,
		"message": message,
		"data": fiber.Map{
			dataName:   data,
			"total":    total,
			"page":     page,
			"per_page": perPage,
		},
	})
}

func RespondError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   true,
		"message": message,
	})
}

func RespondErrorHtmlLogin(c *fiber.Ctx, statusCode int, message string) error {
	csrfToken := c.Locals("csrf").(string)
	return c.Status(statusCode).Render("admin/auth/login", fiber.Map{
		"Error":   true,
		"Token":   csrfToken,
		"Message": message,
	}, "admin/layouts/app")
}

func RespondErrorHtmlDashboard(c *fiber.Ctx, statusCode int, message string, template string) error {
	csrfToken := c.Locals("csrf").(string)
	return c.Status(statusCode).Render(template, fiber.Map{
		"Error":   true,
		"Token":   csrfToken,
		"Message": message,
	}, "admin/template")
}

func RespondErrorHtmlFront(c *fiber.Ctx, statusCode int, message string, template string) error {
	csrfToken := c.Locals("csrf").(string)
	return c.Status(statusCode).Render(template, fiber.Map{
		"Error":   true,
		"Token":   csrfToken,
		"Message": message,
	}, "landing/template")
}
