package routers

import (
	"net/http"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func DeleteParticipation(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Query().Get("eventId")
	if len(eventID) < 1 {
		http.Error(w, "Debe enviar el parámetro eventId", http.StatusBadRequest)
		return
	}
	var t models.Participation
	t.EventID = eventID
	t.UserID = userID

	status, err := db.RemoveParticipation(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al borrar al usuario del evento "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado borrar al usuario del evento "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
