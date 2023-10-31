package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehmet-dogru/go-fiber-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handlers.Hello)
	app.Post("/task", handlers.CreateTask)
	app.Get("/task", handlers.GetTasks)
	app.Get("/task/:id", handlers.GetTaskById)
}
