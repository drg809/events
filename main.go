package main

import (
	"log"

	"github.com/drg809/events/db"
	"github.com/drg809/events/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin connexión a la BD")
		return
	}
	handlers.Handlers()
}
