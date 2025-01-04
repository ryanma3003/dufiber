package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type DashboardController struct {
	authService     service.AuthService
	donationService service.DonationService
	store           *session.Store
}

func NewDashboardController(authService service.AuthService, donationService service.DonationService, store *session.Store) *DashboardController {
	return &DashboardController{authService, donationService, store}
}

func (h *DashboardController) Dashboard(c *fiber.Ctx) error {
	auth := c.Locals("user")
	userSession, ok := auth.(dto.UserSession)
	if !ok {
		return helper.RespondErrorHtmlLogin(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	csrfToken := c.Locals("csrf").(string)

	totalDonatur, err := h.donationService.FindTotalDonatur(c.Context())
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondErrorHtmlDashboard(c, e.Code, e.Message, "admin/pages/dashboard/index")
		}
		return helper.RespondErrorHtmlDashboard(c, fiber.StatusBadRequest, err.Error(), "admin/pages/dashboard/index")
	}

	totalZakat, err := h.donationService.FindTotalZakat(c.Context())
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondErrorHtmlDashboard(c, e.Code, e.Message, "admin/pages/dashboard/index")
		}
		return helper.RespondErrorHtmlDashboard(c, fiber.StatusBadRequest, err.Error(), "admin/pages/dashboard/index")
	}

	totalInfaq, err := h.donationService.FindTotalInfaq(c.Context())
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondErrorHtmlDashboard(c, e.Code, e.Message, "admin/pages/dashboard/index")
		}
		return helper.RespondErrorHtmlDashboard(c, fiber.StatusBadRequest, err.Error(), "admin/pages/dashboard/index")
	}

	totalWakaf, err := h.donationService.FindTotalWakaf(c.Context())
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondErrorHtmlDashboard(c, e.Code, e.Message, "admin/pages/dashboard/index")
		}
		return helper.RespondErrorHtmlDashboard(c, fiber.StatusBadRequest, err.Error(), "admin/pages/dashboard/index")
	}

	fmt.Println(totalDonatur, totalZakat, totalInfaq, totalWakaf)

	return c.Render("admin/pages/dashboard/index", fiber.Map{
		"Auth":         userSession,
		"TotalDonatur": totalDonatur,
		"TotalZakat":   totalZakat,
		"TotalInfaq":   totalInfaq,
		"TotalWakaf":   totalWakaf,
		"Token":        csrfToken}, "admin/template")
}
