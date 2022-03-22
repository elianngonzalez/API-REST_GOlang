package routes

import (
	"API-REST/controller"

	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/personas/api").Subrouter()
	subRoute.HandleFunc("/all", controller.GetAll).Methods("GET")
	subRoute.HandleFunc("/find/{id}", controller.Get).Methods("GET")
	subRoute.HandleFunc("/save", controller.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controller.Delete).Methods("DELETE")
}
