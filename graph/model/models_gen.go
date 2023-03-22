// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/nrfta/go-paging"
)

type Archivo struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

type ArchivoInput struct {
	Ids         []*string `json:"ids,omitempty"`
	Nombre      *string   `json:"nombre,omitempty"`
	ProgramaID  *string   `json:"programaID,omitempty"`
	ProyectoID  *string   `json:"proyectoID,omitempty"`
	ActividadID *string   `json:"actividadID,omitempty"`
	EjecucionID *string   `json:"ejecucionID,omitempty"`
}

type Categoria struct {
	ID      string     `json:"id"`
	Nombre  string     `json:"nombre"`
	Tipo    string     `json:"tipo"`
	PadreID string     `json:"padreID"`
	Padre   *Categoria `json:"padre"`
}

type CategoriaInput struct {
	Ids     []*string `json:"ids,omitempty"`
	Nombre  *string   `json:"nombre,omitempty"`
	Tipo    *string   `json:"tipo,omitempty"`
	PadreID *string   `json:"padreID,omitempty"`
}

type CategoriaNodeConnection struct {
	PageInfo *paging.PageInfo     `json:"pageInfo"`
	Edges    []*CategoriaNodeEdge `json:"edges"`
}

type CategoriaNodeEdge struct {
	Cursor string     `json:"cursor"`
	Node   *Categoria `json:"node,omitempty"`
}

type ChangePasswordInput struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"newPassword"`
}

type Grupo struct {
	ID          string     `json:"id"`
	Nombre      string     `json:"nombre"`
	Descripcion *string    `json:"descripcion,omitempty"`
	Permisos    []*Permiso `json:"permisos,omitempty"`
	Usuarios    []*Usuario `json:"usuarios,omitempty"`
}

type GrupoInput struct {
	Ids         []*string `json:"ids,omitempty"`
	Nombre      *string   `json:"nombre,omitempty"`
	Descripcion *string   `json:"descripcion,omitempty"`
	PermisoID   *string   `json:"permisoID,omitempty"`
	UsuarioID   *string   `json:"usuarioID,omitempty"`
}

type LoginUsuarioInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Permiso struct {
	ID       string     `json:"id"`
	Nombre   string     `json:"nombre"`
	Grupos   []*Grupo   `json:"grupos,omitempty"`
	Usuarios []*Usuario `json:"usuarios,omitempty"`
}

type PermisoInput struct {
	Ids       []*string `json:"ids,omitempty"`
	Nombre    *string   `json:"nombre,omitempty"`
	GrupoID   *string   `json:"grupoID,omitempty"`
	UsuarioID *string   `json:"usuarioID,omitempty"`
}

type Pos struct {
	ID       string `json:"id"`
	Posicion int    `json:"posicion"`
}

type PosInput struct {
	ID       string `json:"id"`
	Posicion int    `json:"posicion"`
}

type RefreshToken struct {
	Token string `json:"token"`
}

type RegistrarUsuarioInput struct {
	Nombres       string `json:"nombres"`
	Apellidos     string `json:"apellidos"`
	NombreUsuario string `json:"nombreUsuario"`
	Email         string `json:"email"`
	Password      string `json:"password"`
}

type SaveCategoriaInput struct {
	Nombre  string  `json:"nombre"`
	Tipo    string  `json:"tipo"`
	PadreID *string `json:"padreID,omitempty"`
}

type SaveGrupoInput struct {
	Nombre      string    `json:"nombre"`
	Descripcion *string   `json:"descripcion,omitempty"`
	PermisosIDs []*string `json:"permisosIDs,omitempty"`
}

type TempToken struct {
	Token string `json:"token"`
}

type UpdateCategoriaInput struct {
	ID      string  `json:"id"`
	Nombre  *string `json:"nombre,omitempty"`
	Tipo    *string `json:"tipo,omitempty"`
	PadreID *string `json:"padreID,omitempty"`
}

type UpdateGrupoInput struct {
	ID          string    `json:"id"`
	Nombre      *string   `json:"nombre,omitempty"`
	Descripcion *string   `json:"descripcion,omitempty"`
	PermisosIDs []*string `json:"permisosIDs,omitempty"`
}

type UpdateUsuarioInput struct {
	ID            string    `json:"id"`
	Status        *bool     `json:"status,omitempty"`
	Nombres       *string   `json:"nombres,omitempty"`
	Apellidos     *string   `json:"apellidos,omitempty"`
	NombreUsuario *string   `json:"nombreUsuario,omitempty"`
	Email         *string   `json:"email,omitempty"`
	GrupoID       *string   `json:"grupoID,omitempty"`
	PermisosIDs   []*string `json:"permisosIDs,omitempty"`
}

type Usuario struct {
	ID            string     `json:"id"`
	Status        bool       `json:"status"`
	Verificado    bool       `json:"verificado"`
	Nombres       string     `json:"nombres"`
	Apellidos     string     `json:"apellidos"`
	NombreUsuario string     `json:"nombreUsuario"`
	Email         string     `json:"email"`
	GrupoID       string     `json:"grupoID"`
	Grupo         *Grupo     `json:"grupo"`
	Permisos      []*Permiso `json:"permisos,omitempty"`
}

type UsuarioInput struct {
	Ids           []*string `json:"ids,omitempty"`
	Status        *bool     `json:"status,omitempty"`
	Verificado    *bool     `json:"verificado,omitempty"`
	Nombres       *string   `json:"nombres,omitempty"`
	Apellidos     *string   `json:"apellidos,omitempty"`
	NombreUsuario *string   `json:"nombreUsuario,omitempty"`
	Email         *string   `json:"email,omitempty"`
	GrupoID       *string   `json:"grupoID,omitempty"`
	PermisosIDs   []*string `json:"permisosIDs,omitempty"`
}
