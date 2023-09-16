package models

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)



type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}


func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	u.Password = string(bytes)
	return nil
}