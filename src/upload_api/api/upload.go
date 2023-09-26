package api

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type upload struct {
	Filename string
	Size float64
	Mimetype string
	Ext string
	Path string
}

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("an error occurred while uploading the file: %s", err))
	}
	fmt.Printf("Filename: %s\n", file.Filename)
	fmt.Printf("Size: %d\n", file.Size)
	fmt.Printf("MIME: %s\n", file.Header["Content-Type"])

	savePath := filepath.Join("./uploads", file.Filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	c.Status(201).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"filename": file.Filename,
	})
	return nil
}