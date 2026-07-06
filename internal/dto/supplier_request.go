package dto

type CreateSupplierRequest struct {
	SupplierName string `json:"supplier_name" validate:"required,min=3,max=100"`
	Nickname     string `json:"nickname" validate:"required"`
	Status       string `json:"status" validate:"required,oneof=Active Inactive"`
}
