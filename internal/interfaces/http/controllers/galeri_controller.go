package controllers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type GaleriController struct {
	galeriService service.GaleriService
}

func NewGaleriController(galeriService service.GaleriService) *GaleriController {
	return &GaleriController{galeriService: galeriService}
}

func (h *GaleriController) GetAllGaleris(c *fiber.Ctx) error {
	per_page, err := strconv.Atoi(c.Query("per_page", "10"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid per_page query value")
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid page query value")
	}

	offset := (page - 1) * per_page

	blogs, err := h.galeriService.FindAllWithPagination(c.Context(), per_page, offset)
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondWithPagination(c, fiber.StatusOK, "success", blogs.TotalData, page, per_page, "donationcategories", blogs.Data)
}

func (h *GaleriController) GetGaleriById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	blog, err := h.galeriService.FindById(c.Context(), id)
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondWithData(c, fiber.StatusOK, "success", blog)
}

func (h *GaleriController) CreateGaleri(c *fiber.Ctx) error {
	userInput := new(dto.GaleriCreate)
	if err := c.BodyParser(userInput); err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := helper.ValidateStruct(userInput); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			// to do add validation input error
			default:
				return helper.RespondError(c, fiber.StatusBadRequest, err.Error())
			}
		}
	}

	if _, err := h.galeriService.Create(c.Context(), userInput); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.RespondMessage(c, fiber.StatusOK, "create success")
}

func (h *GaleriController) EditGaleri(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	userInput := new(dto.GaleriUpdate)
	if err = c.BodyParser(userInput); err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, err.Error())
	}

	userInput.Id = id

	if err = helper.ValidateStruct(userInput); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Id":
				return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
			default:
				return helper.RespondError(c, fiber.StatusBadRequest, err.Error())
			}
		}
	}

	if err = h.galeriService.Update(c.Context(), userInput); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondMessage(c, fiber.StatusOK, "update success")
}

func (h *GaleriController) DeleteGaleri(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	if err = h.galeriService.Delete(c.Context(), id); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.RespondMessage(c, fiber.StatusOK, "delete success")
}
