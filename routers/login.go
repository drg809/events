package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/drg809/events/db"
	"github.com/drg809/events/jwt"
	"github.com/drg809/events/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña incorrectos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 8 {
		http.Error(w, "El password debe contener 8 caráctares cómo mínimo", 400)
		return
	}

	document, exist := db.RetryLogin(t.Email, t.Password)
	if !exist {
		http.Error(w, "Usuario y/o contraseña incorrectos", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el token jwt "+err.Error(), 400)
		return
	}

	resp := models.LoginTokenResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
