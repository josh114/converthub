package convertapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josh114/converthub/src/convert/database"
	"github.com/josh114/converthub/src/convert/models"
)

func GetConverts (c *fiber.Ctx) error {
	converts := []models.Convert{}

	database.Database.Db.Find(&converts)
	responses := []ConvertResponse{}

	for _, response := range converts {
		responseConvert := CreateConvertResponse(response)
		responses = append(responses, responseConvert)
	}
	c.Status(200).JSON(responses)
	return nil
}