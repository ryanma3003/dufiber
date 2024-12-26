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

func (h *FrontController) HomepagePage(c *fiber.Ctx) error {
	homepage, err := h.frontService.HomepageFirst(c.Context())
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/pages/landing", fiber.Map{
		"Homepage": homepage,
	}, "landing/template")
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

func (h *FrontController) ContactPage(c *fiber.Ctx) error {
	contact, err := h.frontService.ContactFirst(c.Context())
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/pages/contact", fiber.Map{
		"Contact": contact,
	}, "landing/template")
}
