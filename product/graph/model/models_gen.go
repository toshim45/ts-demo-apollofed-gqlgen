// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/google/uuid"
)

type Query struct {
}

type Tenant struct {
	ID uuid.UUID `json:"id"`
}

func (Tenant) IsEntity() {}

type Uom struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
