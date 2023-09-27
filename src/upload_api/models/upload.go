package models

import (
	"gorm.io/gorm"
)

type Upload struct {
	gorm.Model
	Filename string 
	Ext string 
	Size int64 
	Path string 
}