package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql/introspection"
)

func (r *__SchemaResolver) Description(ctx context.Context, obj *introspection.Schema) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *__TypeResolver) SpecifiedByURL(ctx context.Context, obj *introspection.Type) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// __Schema returns generated.__SchemaResolver implementation.
func (r *Resolver) __Schema() generated.__SchemaResolver { return &__SchemaResolver{r} }

// __Type returns generated.__TypeResolver implementation.
func (r *Resolver) __Type() generated.__TypeResolver { return &__TypeResolver{r} }

type __SchemaResolver struct{ *Resolver }
type __TypeResolver struct{ *Resolver }
