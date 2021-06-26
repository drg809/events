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

	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")

	router.HandleFunc("/signIn", middlew.CheckDB(routers.SignIn)).Methods("POST")
	router.HandleFunc("/users", middlew.CheckDB(middlew.ValidateJWT(routers.GetUsers))).Methods("GET")
	router.HandleFunc("/users/profile", middlew.CheckDB(middlew.ValidateJWT(routers.UserProfile))).Methods("GET")
	router.HandleFunc("/users/profile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/users/avatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/users/avatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/users/banner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/users/banner", middlew.CheckDB(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/users/follow", middlew.CheckDB(middlew.ValidateJWT(routers.CheckFollow))).Methods("GET")
	router.HandleFunc("/users/follow", middlew.CheckDB(middlew.ValidateJWT(routers.Follow))).Methods("POST")
	router.HandleFunc("/users/unfollow", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteFollow))).Methods("DELETE")

	router.HandleFunc("/events", middlew.CheckDB(middlew.ValidateJWT(routers.GetEvent))).Methods("GET")
	router.HandleFunc("/events", middlew.CheckDB(middlew.ValidateJWT(routers.SaveEvent))).Methods("POST")
	router.HandleFunc("/events", middlew.CheckDB(middlew.ValidateJWT(routers.UpdateEvent))).Methods("PUT")
	router.HandleFunc("/events/photo", middlew.CheckDB(middlew.ValidateJWT(routers.UploadEventPhoto))).Methods("POST")
	router.HandleFunc("/events/photo", middlew.CheckDB(routers.GetEventPhoto)).Methods("GET")
	router.HandleFunc("/events/user", middlew.CheckDB(routers.GetEventsByUserId)).Methods("GET")
	router.HandleFunc("/events", middlew.CheckDB(routers.GetEvents)).Methods("GET")
	router.HandleFunc("/events", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteEvent))).Methods("DELETE")

	router.HandleFunc("/participations", middlew.CheckDB(middlew.ValidateJWT(routers.Participate))).Methods("POST")
	router.HandleFunc("/participations", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteParticipation))).Methods("DELETE")
	router.HandleFunc("/participations/user", middlew.CheckDB(middlew.ValidateJWT(routers.CheckParticipation))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
