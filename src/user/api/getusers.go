package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/user/database"
	"github.com/josh114/converthub/src/user/models"
)


func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
		
	}
	c.Status(200).JSON(responseUsers) 
	return nil
}