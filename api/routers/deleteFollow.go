package routers

import (
	"net/http"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func DeleteFollow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("ID")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	var t models.Follow
	t.UserFollowID = ID
	t.UserID = userID

	status, err := db.RemoveFollow(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al dejar de seguir al usuario "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado dejar de seguir al usuario "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
