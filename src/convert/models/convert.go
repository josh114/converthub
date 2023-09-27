package models

import "gorm.io/gorm"

type Convert struct {
	gorm.Model
	Filename string
	InitialFormat string
	NewFormat string
	Size int64
	Path string
	Guest bool
}