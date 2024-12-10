package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ryanma3003/dufiber/internal/infrastructure/database"
	"github.com/ryanma3003/dufiber/internal/infrastructure/repository"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/controllers"
	"github.com/ryanma3003/dufiber/internal/service"
)

func main() {
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	database.ConnectDB(config)

	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo, database.DB)
	authService := service.NewAuthService(userRepo, database.DB)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Daulat Umat",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		},
	})

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Expiration:        1 * time.Second,
		Max:               10,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   true,
				"message": "Too many requests",
			})
		},
	}))
	app.Use(recover.New())

	app.Static("/", "./public")

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", authController.Login)

	log.Fatal(app.Listen(":3000"))
}
