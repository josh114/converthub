package convertapi

import "github.com/josh114/converthub/src/convert/models"

func SaveConvert(name string, size int64, ext string, path string, format string) (*models.Convert, error) {
	return &models.Convert{
		Filename: name,
		InitialFormat: ext,
		NewFormat: format,
		Size: size,
		Path: path,
		Guest: true,
	}, nil
}