package resolver

import (
	"github.com/sonderkevin/governance/application"
)

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ArchivoService   *application.ArchivoService
	AuthService      *application.AuthService
	CategoriaService *application.CategoriaService
	GrupoService     *application.GrupoService
	PermisoService   *application.PermisoService
	UsuarioService   *application.UsuarioService
}
