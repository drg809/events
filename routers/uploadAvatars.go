package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/drg809/events/db"
	"github.com/drg809/events/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "Debe enviar un avaar ! "+err.Error(), http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var path string = "uploads/avatars/" + userID + "." + extension

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

	var user models.User
	var status bool

	user.Avatar = userID + "." + extension
	status, err = db.ModifyEntry(user, userID)
	if err != nil || !status {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
