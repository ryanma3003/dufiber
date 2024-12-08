package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/ryanma3003/daulatumat/db"
)

func main() {
	config, err := db.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	db.ConnectDB(config)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Daulat Umat",
	})

	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}
