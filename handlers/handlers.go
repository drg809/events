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
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidateJWT(routers.UserProfile))).Methods("GET")
	router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/event", middlew.CheckDB(middlew.ValidateJWT(routers.SaveEvent))).Methods("POST")
	router.HandleFunc("/event", middlew.CheckDB(middlew.ValidateJWT(routers.GetEvents))).Methods("GET")
	router.HandleFunc("/event", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteEvent))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
