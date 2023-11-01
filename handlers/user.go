package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mehmet-dogru/go-fiber-api/database"
	"github.com/mehmet-dogru/go-fiber-api/models"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const secretKey = "secret"

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

func Login(ctx *fiber.Ctx) error {
	db := database.DB
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		panic(err)
	}

	var user models.User

	db.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"message": "Wrong password or email"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func Profile(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return ctx.JSON(user)
}

func Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}
