package jwt

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

var Email string
var UserID string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("MasterdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token incorrecto")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, exist, _ := db.CheckUserExist(claims.Email)
		if exist {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, exist, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token no válido")
	}

	return claims, false, string(""), err
}
