package convertapi

import (
	"time"

	"github.com/josh114/converthub/src/convert/models"
)

type ConvertResponse struct {
ID uint `json:"id"`
CreatedAt time.Time
UpdatedAt time.Time
InitialFormat string `json:"initial_format"`
NewFormat string `json:"format"`
Size int64 `json:"size"`
Filename string `json:"file_name"`

}

func CreateConvertResponse (convert models.Convert) ConvertResponse {
	return ConvertResponse{
		ID: convert.ID, CreatedAt: convert.CreatedAt, UpdatedAt: convert.UpdatedAt, InitialFormat: convert.InitialFormat, NewFormat: convert.NewFormat, Size: convert.Size, Filename: convert.Filename,
	}
}