package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/auth/models"
)

type Token struct {
	Token string `json:"token"`
}

func createResponseToken(t string) Token {
	return Token{Token: t}
}

func Login(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	email := user.Email
	password := user.Password
	
	if !Verify(email, password) {
		return c.Status(401).SendString("invalid email or password")
	}
	token, err := GenerateToken(email)
	if err != nil {
		return c.Status(500).SendString("error generating token, attempt logging in again or contact admin for help")
	}
	// log.Println("this is user data", token)

	c.Status(200).JSON(createResponseToken(token))
	return nil
}