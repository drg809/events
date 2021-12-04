package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/drg809/events/db"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámtro page con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	fmt.Println(page)
	response, result := db.ListEvents(UserID, pag)
	if !result {
		http.Error(w, "Error al leer los eventos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
