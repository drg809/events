package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func Participate(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Query().Get("eventId")
	if len(eventID) < 1 {
		http.Error(w, "Debe enviar el parámetro eventId", http.StatusBadRequest)
		return
	}
	var t models.Participation

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	t.UserID = userID
	t.EventID = eventID

	valid, text := db.CheckTotalParticipants(t)
	if !valid {
		http.Error(w, text, http.StatusBadRequest)
		return
	}

	status, err := db.InsertParticipation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al inscribir al usuario al evento "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado inscribir al usuario al evento "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
