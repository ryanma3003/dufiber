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

	lblog, errl := h.frontService.LastBlog(c.Context())
	if errl != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, errl.Error())
	}

	blogs, errb := h.frontService.BlogAll(c.Context(), 2, 1)
	if errb != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, errb.Error())
	}

	galeris, errg := h.frontService.GaleriAll(c.Context(), 6, 0)
	if errg != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, errg.Error())
	}

	return c.Render("landing/pages/landing", fiber.Map{
		"Homepage": homepage,
		"Blogs":    blogs,
		"Lblog":    lblog,
		"Galeris":  galeris,
	}, "landing/template")
}

func (h *FrontController) AboutPage(c *fiber.Ctx) error {
	about, err := h.frontService.AboutFirst(c.Context())
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/pages/latar", fiber.Map{
		"About": about,
	}, "landing/template")
}

func (h *FrontController) GaleriPage(c *fiber.Ctx) error {
	res, err := h.frontService.GaleriAll(c.Context(), 12, 0)
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/pages/galeri", fiber.Map{
		"Galeris": res,
	}, "landing/template")
}

func (h *FrontController) BlogPage(c *fiber.Ctx) error {
	res, err := h.frontService.BlogAll(c.Context(), 12, 0)
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/article/index_article", fiber.Map{
		"Blogs": res,
	}, "landing/template")
}

func (h *FrontController) BlogShowPage(c *fiber.Ctx) error {
	slug := c.Params("slug")
	res, err := h.frontService.BlogFindBySlug(c.Context(), slug)
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/article/show", fiber.Map{
		"Blog": res,
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

func (h *FrontController) TermPage(c *fiber.Ctx) error {
	term, err := h.frontService.TermFirst(c.Context())
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/pages/term", fiber.Map{
		"Term": term,
	}, "landing/template")
}

func (h *FrontController) PrivacyPage(c *fiber.Ctx) error {
	privacy, err := h.frontService.PrivacyFirst(c.Context())
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Render("landing/pages/privacy", fiber.Map{
		"Privacy": privacy,
	}, "landing/template")
}
