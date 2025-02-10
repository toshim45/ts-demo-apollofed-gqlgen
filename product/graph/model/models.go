package model

import "github.com/google/uuid"

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Sku      string    `json:"sku"`
	Status   string    `json:"status"`
	Type     string    `json:"type"`
	TenantID uuid.UUID `json:"tenant_id"`
	Uom      []Uom     `json:"uom,omitempty"`
}

func (Product) IsEntity() {}
