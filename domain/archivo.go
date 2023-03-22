package domain

import "gorm.io/gorm"

type Archivo struct {
	gorm.Model
	Nombre      string
	ProgramaID  *uint
	ProyectoID  *uint
	ActividadID *uint
	EjecucionID *uint
	PdotID      *uint
	PropMetaID  *uint
}

type ArchivoFilter struct {
	Filter
	Nombre      *string
	ProgramaID  *uint
	ProyectoID  *uint
	ActividadID *uint
	EjecucionID *uint
	PdotID      *uint
	PropMetaID  *uint
}
