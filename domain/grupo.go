package domain

import (
	"gorm.io/gorm"
)

type Grupo struct {
	gorm.Model
	Nombre      string
	Descripcion *string
	Permisos    []*Permiso `gorm:"many2many:grupo_permisos"`
	Usuarios    []*Usuario
}

type GrupoFilter struct {
	Filter
	Nombre      *string
	Descripcion *string
	PermisoID   *uint
	UsuarioID   *uint
}
