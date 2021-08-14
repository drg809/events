package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drg809/events/db"
)

func GetEventParticipants(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("eventId")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro eventId", http.StatusBadRequest)
		return
	}

	response, status := db.GetEventParticipations(ID)
	if status {
		http.Error(w, "Ocurrió un error al buscar el registro ", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
