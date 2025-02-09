package graph

import (
	"context"

	"github.com/google/uuid"
	"github.com/toshim45/demo-apollo-fed-gqlgen/product/graph/model"
)

var (
	uomID1     uuid.UUID = uuid.MustParse("1d56cbd5-e0f1-411b-b7f8-68f7546c2a5e")
	productID1 uuid.UUID = uuid.MustParse("f5713adf-3b1f-4839-9b41-54aec5b69fc9")
	productID2 uuid.UUID = uuid.MustParse("85c34012-4e17-11ee-8739-4201c0a8580d")
	tenantID1  uuid.UUID = uuid.MustParse("8189f6aa-0d53-4362-bf29-f141e5a09438")
	tenantID2  uuid.UUID = uuid.MustParse("b435fdd9-d4c7-499e-9b4f-5910c8d63a75")
)

var UomPiece model.Uom = model.Uom{ID: uomID1, Name: "Piece"}

var DataProductUom map[uuid.UUID][]*model.Uom = map[uuid.UUID][]*model.Uom{
	productID1: {&UomPiece},
	productID2: {&UomPiece},
}

var Data map[uuid.UUID]model.Product = map[uuid.UUID]model.Product{
	productID1: {ID: productID1,
		Name:     "Luna Bright Brightening Serum Concentrate II",
		Sku:      "2243-8899",
		Type:     "bundle",
		TenantID: tenantID1,
	},
	productID2: {ID: productID2,
		Name:     "Oma Elly Pasta Lasagna 250 gram",
		Sku:      "000001T",
		Type:     "inventory",
		TenantID: tenantID2,
	},
}

func FindAll(ctx context.Context) []*model.Product {
	out := []*model.Product{}

	for _, m := range Data {
		if val, exist := DataProductUom[m.ID]; exist {
			m.Uom = val
		}
		out = append(out, &m)
	}

	return out
}

func FindByID(ctx context.Context, id uuid.UUID) *model.Product {
	if val, exist := Data[id]; exist {
		return &val
	}

	return nil
}
