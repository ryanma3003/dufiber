package controllers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/dto"
	"github.com/ryanma3003/dufiber/internal/service"
	"github.com/ryanma3003/dufiber/pkg/helper"
)

type AuthController struct {
	authService service.AuthService
	store       *session.Store
}

func NewAuthController(authService service.AuthService, store *session.Store) *AuthController {
	return &AuthController{authService, store}
}

func (h *AuthController) LoginPage(c *fiber.Ctx) error {
	// check if user is authenticated
	sess, err := h.store.Get(c)
	if err == nil {
		userID := sess.Get("user_id")
		if userID != nil {
			// Redirect to dashboard if session is still active
			return c.Redirect("/duadmin/dashboard")
		}
	}

	// set csrf token
	csrfToken := c.Locals("csrf").(string)

	// render html
	return c.Render("admin/auth/login", fiber.Map{
		"Token": csrfToken,
	}, "admin/layouts/app")
}

func (h *AuthController) Login(c *fiber.Ctx) error {
	// initiate input struct
	loginInput := new(dto.LoginInput)

	// get the input value
	username := c.FormValue("username")
	password := c.FormValue("password")

	// populate struct with value
	loginInput.Username = username
	loginInput.Password = password

	// bind value to struct
	if err := c.BodyParser(loginInput); err != nil {
		return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, err.Error())
	}

	// validate input
	if err := helper.ValidateStruct(loginInput); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Username":
				// return helper.RespondError(c, fiber.StatusBadRequest, "Invalid username")
				return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, "Username is required")
			case "Password":
				return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, "Password is required")
			default:
				return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, err.Error())
			}
		}

	}

	// check if user exist
	user, err := h.authService.LoginUser(c.Context(), loginInput)
	if err != nil {
		if e, ok := err.(helper.AppError); ok {
			return helper.RespondErrorHtmlLogin(c, e.Code, e.Message)
		}
		return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, err.Error())
	}

	// Create a session and store user information
	sess, err := h.store.Get(c)
	if err != nil {
		return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, err.Error())
	}

	// set session variable
	sess.Set("user_id", user.Id)
	sess.Set("role", user.Role)
	sess.Set("username", user.Username)
	sess.Set("name", user.Name)

	// save session
	if err := sess.Save(); err != nil {
		return helper.RespondErrorHtmlLogin(c, fiber.StatusBadRequest, err.Error())
	}

	fmt.Println(sess.Get("user_id"))
	fmt.Println(sess.Get("role"))
	fmt.Println(sess.Get("username"))
	fmt.Println(sess.Get("name"))

	// return for API with token
	// return helper.RespondWithData(c, fiber.StatusOK, "Login success", fiber.Map{
	// 	"token":        token.Token,
	// 	"token_type":   "Bearer",
	// 	"expired_time": "8h",
	// })

	// for debugging
	// return helper.RespondWithData(c, fiber.StatusOK, "Login success", fiber.Map{
	// 	"user": user,
	// })

	// redirect to dashboard
	return c.Redirect("/duadmin/dashboard")
}

func (h *AuthController) Logout(c *fiber.Ctx) error {
	// Get the session
	sess, err := h.store.Get(c)
	if err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, "Failed to get session")
	}

	// Destroy the session
	if err := sess.Destroy(); err != nil {
		return helper.RespondError(c, fiber.StatusInternalServerError, "Failed to destroy session")
	}

	// Clear any cache if necessary (example, if using a cache service)
	// cacheService.ClearCache(userID)

	return c.Redirect("/duadmin/login")
}
