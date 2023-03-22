package domain

import "gorm.io/gorm"

type Permiso struct {
	gorm.Model
	Nombre   string
	Grupos   []*Grupo   `gorm:"many2many:grupo_permisos"`
	Usuarios []*Usuario `gorm:"many2many:usuario_permisos"`
}

type PermisoFilter struct {
	Filter
	Nombre      *string
	GruposIDs   *[]uint
	UsuariosIDs *[]uint
}
