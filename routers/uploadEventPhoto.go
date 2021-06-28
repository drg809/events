package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func UploadEventPhoto(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Debe enviar una foto ! "+err.Error(), http.StatusBadRequest)
		return
	}
	eventID := r.URL.Query().Get("eventId")
	if len(eventID) < 1 {
		http.Error(w, "Debe enviar el parámetro eventID", http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var path string = "uploads/events/" + eventID + "." + extension

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var event models.Event
	var status bool

	event.Photo = eventID + "." + extension
	status, err = db.UpdateEvent(event, eventID)
	fmt.Println(status)
	fmt.Println(event.Photo)
	if err != nil || !status {
		http.Error(w, "Error al actualizar el evento ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
