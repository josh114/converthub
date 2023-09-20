package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/user/database"
	"github.com/josh114/converthub/src/user/models"
)

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}
func GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Status(400).JSON("Please ensure that :id is an integer")
	}
	var user models.User
	if err := findUser(id, &user); err != nil {
		c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	c.Status(200).JSON(responseUser)
	return nil
}