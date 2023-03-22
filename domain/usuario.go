package domain

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Status             bool
	Verificado         bool
	CodigoVerificacion string
	Nombres            string
	Apellidos          string
	NombreUsuario      string
	FotoPerfil         string
	Email              string
	Password           string
	GrupoID            uint
	Permisos           []*Permiso   `gorm:"many2many:usuario_permisos"`
}

type UsuarioFilter struct {
	Filter
	Status         *bool
	Verificado     *bool
	Nombres        *string
	Apellidos      *string
	NombreUsuario  *string
	FotoPerfil     *string
	Email          *string
	GrupoID        *uint
	PermisosIDs    *[]uint
}
