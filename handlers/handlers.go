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
	router.HandleFunc("/events", middlew.CheckDB(middlew.ValidateJWT(routers.SaveEvent))).Methods("POST")
	router.HandleFunc("/events/user", middlew.CheckDB(routers.GetEventsByUserId)).Methods("GET")
	router.HandleFunc("/events", middlew.CheckDB(routers.GetEvents)).Methods("GET")
	router.HandleFunc("/events", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteEvent))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDB(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.CheckDB(middlew.ValidateJWT(routers.GetBanner))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
