package main

import (
	"log"
	"strings"
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

// var viewsfs embed.FS

func Nl2brHtml(value interface{}) string {
	if str, ok := value.(string); ok {
		return strings.Replace(str, "\n", "<br />", -1)
	}
	return ""
}

func inc(i int) int {
	return i + 1
}

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
	frontRepo := repository.NewFrontRepository()
	// blogRepo := repository.NewBlogRepository()
	// blogCategoryRepo := repository.NewBlogCategoryRepository()

	// service init
	// userService := service.NewUserService(userRepo, database.DB)
	authService := service.NewAuthService(userRepo, database.DB)
	frontService := service.NewFrontService(frontRepo, database.DB)
	// blogService := service.NewBlogService(blogRepo, database.DB)
	// blogCategoryService := service.NewBlogCategoryService(blogCategoryRepo, database.DB)

	// controller init
	// blogController := controllers.NewBlogController(blogService)
	// blogCategoryController := controllers.NewBlogCategoryController(blogCategoryService)
	// userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	frontController := controllers.NewFrontController(frontService)

	// engine html init
	// engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	// engine := django.NewPathForwardingFileSystem(http.FS(viewsfs), "/views", ".django")
	engine := html.New("./views", ".html")
	engine.AddFunc("nl2br", Nl2brHtml)
	engine.AddFunc("inc", inc)

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
	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

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
	app.Get("/", frontController.HomepagePage)
	app.Get("/hubungi-kami", frontController.ContactPage)
	app.Get("/faq", frontController.FaqPage)

	// backend route

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", authController.Login)

	// server run
	log.Fatal(app.Listen(":3000"))
}
