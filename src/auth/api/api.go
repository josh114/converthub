package api

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/auth/database"
	"github.com/josh114/converthub/src/auth/models"
)

func findUser(email string, user *models.User) error {
	database.Database.Db.Find(&user, "email = ?", email)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func Login(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(400).JSON(err.Error())
	}
	email := user.Email
	if err := findUser(email, &user); err != nil {
		c.Status(400).JSON(err.Error())
	}
	log.Println("this is user data", user)

	c.Status(200).SendString("auth data accepted")
	return nil
}