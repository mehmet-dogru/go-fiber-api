package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehmet-dogru/go-fiber-api/database"
	"github.com/mehmet-dogru/go-fiber-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	db := database.DB
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		panic(err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	db.Create(&user)
	return ctx.JSON(user)
}
