package models

import "github.com/jinzhu/gorm"

type Persona struct {
	gorm.Model

	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Dni      string `json:"dni" `
	Telefono string `json:"telefono"`
}
