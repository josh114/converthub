package uploadapi

import (
	"github.com/josh114/converthub/src/convert/database"
	"github.com/josh114/converthub/src/convert/models"
)


func UploadDB(data *models.Upload) error {
	err := database.Database.Db.Create(&data)
	if err != nil {
		return err.Error
	}
	return nil
}