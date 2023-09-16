package database

import (
	"log"
	"os"

	"github.com/josh114/converthub/src/user/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB 
}

var Database DbInstance

func ConnectDb () {
	dburl := "media_user:media1234@tcp(127.0.0.1:3306)/media?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err.Error())
		os.Exit(2 )
	}
	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	db.AutoMigrate(&models.User{} )

	Database = DbInstance{Db: db}
	
}