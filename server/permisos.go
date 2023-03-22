package server

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sonderkevin/governance/graph/resolver"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func RootPermisoMiddleware(ctx context.Context, next graphql.RootResolver) graphql.Marshaler {
	fc := graphql.GetRootFieldContext(ctx)
	if fc == nil {
		return nil
	}

	if !resolver.IsAuthorized(ctx, fc.Field.Name) {
		err := gqlerror.Error{
			Message: "acces denied",
		}
		graphql.AddError(ctx, &err)
		return graphql.MarshalAny(nil)
	}

	return next(ctx)
}
