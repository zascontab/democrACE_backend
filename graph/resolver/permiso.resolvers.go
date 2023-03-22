package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"

	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/governance/graph/generated"
	"github.com/sonderkevin/governance/graph/model"
)

// Grupos is the resolver for the grupos field.
func (r *permisoResolver) Grupos(ctx context.Context, obj *model.Permiso, input *model.GrupoInput, page *paging.PageArgs) ([]*model.Grupo, error) {
	return r.GrupoService.GetAll(&model.GrupoInput{PermisoID: &obj.ID})
}

// Usuarios is the resolver for the usuarios field.
func (r *permisoResolver) Usuarios(ctx context.Context, obj *model.Permiso, input *model.UsuarioInput, page *paging.PageArgs) ([]*model.Usuario, error) {
	return r.UsuarioService.GetAll(&model.UsuarioInput{PermisosIDs: []*string{&obj.ID}})
}

// Permiso returns generated.PermisoResolver implementation.
func (r *Resolver) Permiso() generated.PermisoResolver { return &permisoResolver{r} }

type permisoResolver struct{ *Resolver }