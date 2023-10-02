package uploadapi

import "github.com/josh114/converthub/src/convert/models"

func SaveUpload(name string, size int64, ext string, path string) (*models.Upload, error) {
	return &models.Upload{
		Filename: name,
		Ext: ext,
		Size: size,
		Path: path,
		Service: "convert",
	}, nil
}