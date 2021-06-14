package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/drg809/events/db"
)

func GetEventsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro page cómo un entero mayor a 0", http.StatusBadRequest)
		return
	}

	resp, status := db.ListEventsFollowers(userID, page)
	if !status {
		http.Error(w, "Error al leer los eventos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
