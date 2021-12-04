package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var t models.Event

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	if t.UserID != UserID {
		http.Error(w, "No puede borrar un evento que no es suyo, el incidente será reportado."+err.Error(), http.StatusUnauthorized)
		return
	}

	var status bool

	status, err = db.UpdateEvent(t, UserID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusOK)
}
