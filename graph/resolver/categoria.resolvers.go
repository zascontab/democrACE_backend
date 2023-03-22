package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"

	"github.com/sonderkevin/governance/graph/generated"
	"github.com/sonderkevin/governance/graph/model"
)

// Padre is the resolver for the padre field.
func (r *categoriaResolver) Padre(ctx context.Context, obj *model.Categoria) (*model.Categoria, error) {
	return r.CategoriaService.Get(&obj.PadreID)
}

// Categoria returns generated.CategoriaResolver implementation.
func (r *Resolver) Categoria() generated.CategoriaResolver { return &categoriaResolver{r} }

type categoriaResolver struct{ *Resolver }
