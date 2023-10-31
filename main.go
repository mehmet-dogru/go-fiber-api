package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mehmet-dogru/go-fiber-api/database"
	"github.com/mehmet-dogru/go-fiber-api/router"
	"log"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
