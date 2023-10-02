package convertapi

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func StartConvert(c *fiber.Ctx) error {
	var body Body
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("error occured while parsing body: %s", err))
	}
	name := body.Filename
	ext := body.Format
	input := fmt.Sprintf("./uploads/%s", name)
	title, err := ExtractTitle(name)
	if err != nil {
		fmt.Println("err occured")
	}

	if !FileExists(input) {
		c.Status(404).SendString("the file does not exist. please re-upload file")
	}
	output := fmt.Sprintf("./converts/%s.%s", title, body.Format)
	errr := Convert(input, output, ext)
	if errr != nil {
		return c.Status(400).SendString(fmt.Sprintf("Transcoding failed: %v", err.Error()))
	}
	size, err := os.Stat(output)
	if err != nil {
		log.Println("cannot find file")
	}

	convert, err := SaveConvert(name, size.Size(), ext, output, ext)
	if err != nil {
		log.Println(err.Error())
	}
	go ConvertDB(convert)

	c.Status(200).JSON(fiber.Map{
		"filename": fmt.Sprintf(title, ext),
		"file conversion": "successful",
	})
	return nil
}