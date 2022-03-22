package models

import "github.com/jinzhu/gorm"

type Persona struct {
	gorm.Model
	IdPersona int64  `json:"idPersona"` // gorm:"primary_key;auto_increment"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Dni       string `json:"dni" gorm:"unique"`
	Telefono  string `json:"telefono"`
}
