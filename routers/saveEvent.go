package routers

import (
	"encoding/json"
	"net/http"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func SaveEvent(w http.ResponseWriter, r *http.Request) {

	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Error al parsear la información "+err.Error(), 400)
		return
	}

	entry := models.Event{
		UserID: UserID,
		Name:   event.Name,
		Detail: event.Detail,
		Date:   event.Date,
		Type:   event.Type,
		Photo:  event.Photo,
	}

	_, status, err := db.InsertEvent(entry)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar un nuevo evento, reintente de nuevo "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar el evento", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
