package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehmet-dogru/go-fiber-api/database"
	"github.com/mehmet-dogru/go-fiber-api/models"
)

func Hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello")
}

func GetTasks(ctx *fiber.Ctx) error {
	db := database.DB
	var tasks []models.Task

	db.Find(&tasks)
	return ctx.JSON(tasks)
}

func CreateTask(ctx *fiber.Ctx) error {
	db := database.DB
	task := new(models.Task)

	if err := ctx.BodyParser(task); err != nil {
		return ctx.Status(400).JSON(err)
	}

	db.Create(&task)
	return ctx.JSON(task)
}

func GetTaskById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DB
	var task models.Task
	db.Find(&task, id)

	return ctx.JSON(task)
}
