package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/drg809/events/db"
)

func GetEventsByUserId(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámtro page con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	response, result := db.ListEventsByUserId(ID, int64(page))
	if !result {
		http.Error(w, "Error al leer los eventosu", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
