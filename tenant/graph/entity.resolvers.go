package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"

	"github.com/google/uuid"
	"github.com/toshim45/demo-apollo-fed-gqlgen/tenant/graph/model"
)

// FindTenantByID is the resolver for the findTenantByID field.
func (r *entityResolver) FindTenantByID(ctx context.Context, id uuid.UUID) (model.Tenant, error) {
	return FindByID(ctx, id), nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
