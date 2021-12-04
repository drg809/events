package user

import (
	"errors"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func GetUsers(UserID string, pag int64, search, typeUser string) ([]*models.ListUsers, error) {
	result, status := db.ListUsers(UserID, pag, search, typeUser)
	if !status {
		return nil, errors.New("Error al leer los usuarios")
	}
	return result, nil
}
