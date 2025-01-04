package main

import (
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html/v2"
	"github.com/ryanma3003/dufiber/internal/infrastructure/database"
	"github.com/ryanma3003/dufiber/internal/infrastructure/repository"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/controllers"
	"github.com/ryanma3003/dufiber/internal/interfaces/http/middleware"
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

func unescape(s string) template.HTML {
	return template.HTML(s)
}

func main() {
	// load env
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// connect to database
	database.ConnectDB(config)

	// session
	store := session.New()

	// repo init
	userRepo := repository.NewUserRepository()
	frontRepo := repository.NewFrontRepository()
	donationRepo := repository.NewDonationRepository()
	// blogRepo := repository.NewBlogRepository()
	// blogCategoryRepo := repository.NewBlogCategoryRepository()

	// service init
	// userService := service.NewUserService(userRepo, database.DB)
	authService := service.NewAuthService(userRepo, database.DB)
	frontService := service.NewFrontService(frontRepo, database.DB)
	donationService := service.NewDonationService(donationRepo, database.DB)
	// blogService := service.NewBlogService(blogRepo, database.DB)
	// blogCategoryService := service.NewBlogCategoryService(blogCategoryRepo, database.DB)

	// controller init
	// blogController := controllers.NewBlogController(blogService)
	// blogCategoryController := controllers.NewBlogCategoryController(blogCategoryService)
	// userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, store)
	dashboardController := controllers.NewDashboardController(authService, donationService, store)
	frontController := controllers.NewFrontController(frontService)

	// engine html init
	// engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	// engine := django.NewPathForwardingFileSystem(http.FS(viewsfs), "/views", ".django")
	engine := html.New("./views", ".html")
	engine.AddFunc("nl2br", Nl2brHtml)
	engine.AddFunc("inc", inc)
	engine.AddFunc("unescape", unescape)

	// fiber app init
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Daulat Umat",
		Views:        engine,
	})

	// static file
	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	// middleware
	app.Use(cors.New())
	app.Use(csrf.New(csrf.Config{
		KeyLookup:    "form:_csrf",
		ContextKey:   "csrf",
		Expiration:   1 * time.Hour,
		KeyGenerator: utils.UUIDv4,
	}))

	// encrypt cookie
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key:    config.EncryptCookie,
		Except: []string{"__Host-csrf_", "csrf_", "session_id"}, // exclude CSRF cookie
	}))

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
	app.Get("/tentang-kami", frontController.AboutPage)
	app.Get("/galeri", frontController.GaleriPage)
	app.Get("/artikel", frontController.BlogPage)
	app.Get("/artikel/:slug", frontController.BlogShowPage)
	app.Get("/hubungi-kami", frontController.ContactPage)
	app.Get("/faq", frontController.FaqPage)
	app.Get("/syarat-ketentuan", frontController.TermPage)
	app.Get("/kebijakan-privasi", frontController.PrivacyPage)

	// backend route

	backend := app.Group("/duadmin")

	backend.Get("/login", authController.LoginPage)
	backend.Post("/login", authController.Login)

	backend.Get("/dashboard", middleware.WebAuth(store), dashboardController.Dashboard)

	// logout
	backend.Post("/logout", middleware.WebAuth(store), authController.Logout)

	// v1 := api.Group("/v1")

	// server run
	log.Fatal(app.Listen(":3000"))
}
