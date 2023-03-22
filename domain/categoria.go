package domain

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model
	Nombre  string
	Tipo    string
	PadreID uint
	Padre   *Categoria `gorm:"many2one:PadreID"`
}

type CategoriaFilter struct {
	Filter
	Nombre  *string
	Tipo    *string
	PadreID *uint
}

type UpdateCategoriaInput struct {
	Nombre  *string
	Tipo    *string
	PadreID *uint
}
