package controller

import (
	"API-REST/commons"
	"API-REST/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConection()
	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)

	commons.SendResponse(w, http.StatusOK, json)
}

func Get(w http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]
	db := commons.GetConection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendResponse(w, http.StatusOK, json)
	} else {
		commons.SendError(w, http.StatusBadRequest)
	}

}

func Save(w http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := commons.GetConection()
	defer db.Close()

	err := json.NewDecoder(request.Body).Decode(&persona)
	if err != nil {
		commons.SendError(w, http.StatusBadRequest)
		return
	}

	err = db.Save(&persona).Error //con db.Save crea o actualiza el registro
	if err != nil {
		commons.SendError(w, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)
	commons.SendResponse(w, http.StatusOK, json)

}

func Delete(w http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]
	db := commons.GetConection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		err := db.Delete(&persona).Error
		if err != nil {
			commons.SendError(w, http.StatusInternalServerError)
			return
		}
		commons.SendResponse(w, http.StatusOK, []byte(""))
	} else {
		commons.SendError(w, http.StatusNotFound)
	}

}
