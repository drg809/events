package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drg809/events/db"
)

func GetEvent(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	response, err := db.GetEvent(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
