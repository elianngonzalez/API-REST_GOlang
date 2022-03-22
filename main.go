package main

import (
	"API-REST/commons"
	"API-REST/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)
	commons.EnableCORS(router)

	puerto := ":8000"
	server := http.Server{
		Addr:    puerto,
		Handler: router,
	}

	log.Println("servidor ejecutable en el puerto " + puerto)
	log.Println(server.ListenAndServe())

}
