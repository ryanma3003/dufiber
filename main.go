package main

import (
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/ryanma3003/dufiber/internal/infrastructure/database"
	"github.com/ryanma3003/dufiber/internal/infrastructure/repository"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/controllers"
	"github.com/ryanma3003/dufiber/internal/service"
)

var viewsfs embed.FS

func main() {
	// load env
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// connect to database
	database.ConnectDB(config)

	// repo init
	userRepo := repository.NewUserRepository()
	blogRepo := repository.NewBlogRepository()
	blogCategoryRepo := repository.NewBlogCategoryRepository()

	// service init
	userService := service.NewUserService(userRepo, database.DB)
	authService := service.NewAuthService(userRepo, database.DB)
	blogService := service.NewBlogService(blogRepo, database.DB)
	blogCategoryService := service.NewBlogCategoryService(blogCategoryRepo, database.DB)

	// controller init
	blogController := controllers.NewBlogController(blogService)
	blogCategoryController := controllers.NewBlogCategoryController(blogCategoryService)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)

	// engine html init
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	// fiber app init
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Daulat Umat",
		Views:         engine,
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

	// static file
	app.Static("/", "./public")

	// middleware
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

	// frontend route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/landing/index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	// backend route

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", authController.Login)

	// server run
	log.Fatal(app.Listen(":3000"))
}
