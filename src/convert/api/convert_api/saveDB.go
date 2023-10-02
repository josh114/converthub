package convertapi

import (
	"github.com/josh114/converthub/src/convert/database"
	"github.com/josh114/converthub/src/convert/models"
)


func ConvertDB(data *models.Convert) error {
	err := database.Database.Db.Create(&data)
	if err != nil {
		return err.Error
	}
	return nil
}