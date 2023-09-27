package api

import (
	"github.com/josh114/converthub/src/upload_api/database"
	"github.com/josh114/converthub/src/upload_api/models"
)


func UploadDB(data *models.Upload) error {
	err := database.Database.Db.Create(&data)
	if err != nil {
		return err.Error
	}
	return nil
}