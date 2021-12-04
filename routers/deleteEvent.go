package routers

import (
	"net/http"

	"github.com/drg809/events/db"
)

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := db.DeleteEvent(ID, UserID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el evento "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
