package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func CheckParticipation(w http.ResponseWriter, r *http.Request) {
	eventID := r.URL.Query().Get("eventId")
	if len(eventID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	var t models.Participation
	t.EventID = eventID
	t.UserID = UserID

	var resp models.CheckParticipation

	status, err := db.CheckParticipation(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		valid, _ := db.CheckTotalParticipants(t)
		if valid {
			resp.Status = true
		} else {
			resp.Status = false
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
