package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func Follow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("ID")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	var t models.Follow

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	t.UserID = userID
	t.UserFollowID = ID

	status, err := db.InsertFollow(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al seguir al usuario "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado seguir al usuario "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
