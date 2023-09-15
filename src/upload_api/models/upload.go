package models

import "time"

type Upload struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Filename string `json:"filename"`
	Ext string `json:"ext"`
	Mimetype string `json:"mimetype"`
	Path string `json:"filepath"`
}