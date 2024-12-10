package controllers

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type BlogController struct {
	blogService service.BlogService
}

func NewBlogController(blogService service.BlogService) *BlogController {
	return &BlogController{blogService: blogService}
}

func (h *BlogController) GetAllBlogs(c *fiber.Ctx) error {
	per_page, err := strconv.Atoi(c.Query("per_page", "10"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid per_page query value")
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid page query value")
	}

	offset := (page - 1) * per_page

	blogs, err := h.blogService.FindAllWithPagination(c.Context(), per_page, offset)
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondWithPagination(c, fiber.StatusOK, "success get blogs data", blogs.TotalData, page, per_page, "blogs", blogs.Data)
}

func (h *BlogController) GetBlogById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	blog, err := h.blogService.FindById(c.Context(), id)
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondWithData(c, fiber.StatusOK, "success get blog data", blog)
}

func (h *BlogController) CreateBlog(c *fiber.Ctx) error {
	userInput := new(dto.BlogCreate)
	if err := c.BodyParser(userInput); err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid input")
	}

	if err := helper.ValidateStruct(userInput); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				return helper.RespondError(c, fiber.StatusBadRequest, "Required")
			case "BlogCategoryId":
				return helper.RespondError(c, fiber.StatusBadRequest, "Required")
			default:
				return helper.RespondError(c, fiber.StatusBadRequest, err.Error())
			}
		}
	}

	if _, err := h.blogService.Create(c.Context(), userInput); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.RespondMessage(c, fiber.StatusOK, "success create blog")
}

func (h *BlogController) EditBlog(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	userInput := new(dto.BlogUpdate)
	if err = c.BodyParser(userInput); err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, err.Error())
	}

	userInput.Id = id

	if err = helper.ValidateStruct(userInput); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Id":
				return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
			case "BlogCategoryId":
				return helper.RespondError(c, fiber.StatusBadRequest, "Required")
			default:
				return helper.RespondError(c, fiber.StatusBadRequest, err.Error())
			}
		}
	}

	if err = h.blogService.Update(c.Context(), userInput); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}
	return helper.RespondMessage(c, fiber.StatusOK, "success update blog")
}

func (h *BlogController) DeleteBlog(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helper.RespondError(c, fiber.StatusBadRequest, "Invalid ID")
	}

	if err = h.blogService.Delete(c.Context(), id); err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondError(c, e.Code, e.Message)
		}
		return helper.RespondError(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.RespondMessage(c, fiber.StatusOK, "success delete blog")
}
