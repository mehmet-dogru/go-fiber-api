package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehmet-dogru/go-fiber-api/handlers"
)

func SetupRoutes(app *fiber.App) {

	//TASK ROUTES
	app.Get("/hello", handlers.Hello)
	app.Post("/task", handlers.CreateTask)
	app.Get("/task", handlers.GetTasks)
	app.Get("/task/:id", handlers.GetTaskById)
	app.Put("/task/:id", handlers.UpdateTask)
	app.Delete("/task/:id", handlers.DeleteTask)

	//USER ROUTES
	app.Post("/register", handlers.Register)
}
