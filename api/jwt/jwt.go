package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drg809/events/models"
)

func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("MasterdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":    t.Email,
		"name":     t.Name,
		"surname":  t.Surname,
		"date":     t.Date,
		"bio":      t.Bio,
		"location": t.Location,
		"web":      t.Web,
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
