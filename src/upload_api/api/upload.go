package api

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type Upload struct {
	ID uint `json:"id"`
	Filename string `json:"filename"`
	Size int `json:"size"`
	Mimetype string `json:"mime_type"`
	Ext string `json:"ext"`
	Path string `json:"path"`
}



func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("an error occurred while uploading the file: %s", err))
	}

	s := file.Header["Content-Type"]
	ext := ReturnMime(s[0])

	savePath := filepath.Join("./uploads", file.Filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}
	
	upload, err := SaveUpload(file.Filename, file.Size, ext, savePath)
	if err != nil {
		c.Status(500).SendString(fmt.Sprintf("could not save data to database due to internal error: %s", err.Error()))
	}
	go UploadDB(upload)
	c.Status(201).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"filename": file.Filename,
	})
	return nil
}