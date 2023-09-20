package api

import (
	"errors"
	"log"

	"github.com/josh114/converthub/src/auth/database"
	"github.com/josh114/converthub/src/auth/models"
	"golang.org/x/crypto/bcrypt"
)

func findUser(email string, user *models.User) error {
	database.Database.Db.Find(&user, "email = ?", email)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func Verify (email string, password string) bool {
	var user models.User
	if err := findUser(email, &user); err != nil {
		log.Println(err.Error())
	}
	hash := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}