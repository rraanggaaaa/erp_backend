package model

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	ID           uuid.UUID `json:"id"`
	SupplierCode string    `json:"supplier_code"`
	SupplierName string    `json:"supplier_name"`
	Nickname     string    `json:"nickname"`
	LogoURL      *string   `json:"logo_url"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
