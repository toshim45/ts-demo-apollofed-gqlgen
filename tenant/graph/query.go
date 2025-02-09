package graph

import (
	"context"

	"github.com/google/uuid"
	"github.com/toshim45/demo-apollo-fed-gqlgen/tenant/graph/model"
)

var (
	tenantID1 uuid.UUID = uuid.MustParse("8189f6aa-0d53-4362-bf29-f141e5a09438")
	tenantID2 uuid.UUID = uuid.MustParse("b435fdd9-d4c7-499e-9b4f-5910c8d63a75")
)

var Data map[uuid.UUID]model.Tenant = map[uuid.UUID]model.Tenant{
	tenantID1: {ID: tenantID1, Name: "Jafra", Number: "JFR", Active: true},
	tenantID2: {ID: tenantID2, Name: "Tokonow", Number: "TKNW"},
}

func FindAll(ctx context.Context) []*model.Tenant {
	out := []*model.Tenant{}

	for _, m := range Data {
		out = append(out, &m)
	}

	return out
}

func FindByID(ctx context.Context, id uuid.UUID) *model.Tenant {
	if val, exist := Data[id]; exist {
		return &val
	}

	return nil
}
