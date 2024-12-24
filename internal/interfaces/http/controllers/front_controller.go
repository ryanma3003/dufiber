package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type FrontController struct {
	frontService service.FrontService
}

func NewFrontController(frontService service.FrontService) *FrontController {
	return &FrontController{frontService: frontService}
}

func (h *FrontController) HomepageFirst(c *fiber.Ctx) error {
	res, err := h.frontService.HomepageFirst(c.Context())
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondWithData(c, fiber.StatusOK, "success", res)
}

func (h *FrontController) FaqPage(c *fiber.Ctx) error {
	res, err := h.frontService.FaqAll(c.Context())
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/pages/faq", fiber.Map{
		"Faqs": res,
	}, "landing/template")
}
