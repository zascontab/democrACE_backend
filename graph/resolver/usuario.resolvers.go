package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"

	"github.com/sonderkevin/governance/graph/generated"
	"github.com/sonderkevin/governance/graph/model"
)

// Grupo is the resolver for the grupo field.
func (r *usuarioResolver) Grupo(ctx context.Context, obj *model.Usuario) (*model.Grupo, error) {
	return r.GrupoService.Get(&obj.GrupoID)
}

// Permisos is the resolver for the permisos field.
func (r *usuarioResolver) Permisos(ctx context.Context, obj *model.Usuario) ([]*model.Permiso, error) {
	return r.PermisoService.GetAll(&model.PermisoInput{UsuarioID: &obj.ID})
}

// Usuario returns generated.UsuarioResolver implementation.
func (r *Resolver) Usuario() generated.UsuarioResolver { return &usuarioResolver{r} }

type usuarioResolver struct{ *Resolver }
