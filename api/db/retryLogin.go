package db

import (
	"github.com/drg809/events/models"
	"golang.org/x/crypto/bcrypt"
)

func RetryLogin(email string, password string) (models.User, bool) {
	user, exist, _ := CheckUserExist(email)
	if !exist {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
