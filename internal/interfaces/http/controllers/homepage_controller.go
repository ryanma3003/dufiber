package controllers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type HomepageController struct {
	homepageService service.HomepageService
}

func NewHomepageController(homepageService service.HomepageService) *HomepageController {
	return &HomepageController{homepageService: homepageService}
}

func (h *HomepageController) GetAllHomepages(c *fiber.Ctx) error {
	per_page, err := strconv.Atoi(c.Query("per_page", "10"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid per_page query value")
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid page query value")
	}

	offset := (page - 1) * per_page

	blogs, err := h.homepageService.FindAllWithPagination(c.Context(), per_page, offset)
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondWithPagination(c, fiber.StatusOK, "success", blogs.TotalData, page, per_page, "donationcategories", blogs.Data)
}

func (h *HomepageController) GetHomepageById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	blog, err := h.homepageService.FindById(c.Context(), id)
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondWithData(c, fiber.StatusOK, "success", blog)
}

func (h *HomepageController) CreateHomepage(c *fiber.Ctx) error {
	userInput := new(dto.HomepageCreate)
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

	if _, err := h.homepageService.Create(c.Context(), userInput); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.RespondMessage(c, fiber.StatusOK, "create success")
}

func (h *HomepageController) EditHomepage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	userInput := new(dto.HomepageUpdate)
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

	if err = h.homepageService.Update(c.Context(), userInput); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondMessage(c, fiber.StatusOK, "update success")
}

func (h *HomepageController) DeleteHomepage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	if err = h.homepageService.Delete(c.Context(), id); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.RespondMessage(c, fiber.StatusOK, "delete success")
}
