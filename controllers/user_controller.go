package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/drg809/events/jwt"
	"github.com/drg809/events/models/user"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	userList, err := user.GetUsers(jwt.UserID, pag, search, typeUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SendOKHttp(w, userList)
}

func SendOKHttp(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
