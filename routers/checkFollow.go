package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func CheckFollow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	var t models.Follow
	t.UserFollowID = ID
	t.UserID = UserID

	var resp models.CheckParticipation

	status, err := db.CheckFollow(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
