package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/user/database"
	"github.com/josh114/converthub/src/user/models"
)

type User struct {
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Username string `json:"username"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName, Email: userModel.Email, Username: userModel.Username }
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
	 	c.Status(400).JSON(err.Error())
	}
	log.Println("this is user data", user)
	if err:= user.HashPassword(); err != nil {
		c.Status(400).JSON(err.Error())
	}
	log.Println("this is user data", user)
	database.Database.Db.Create(user)
	responseUser := CreateResponseUser(user)
	c.Status(200).JSON(responseUser)

	return nil
}

