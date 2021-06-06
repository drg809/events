package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/drg809/events/middlew"
	"github.com/drg809/events/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signIn", middlew.CheckDB(routers.SignIn)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
