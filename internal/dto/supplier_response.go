package dto

type SupplierResponse struct {
	ID           string `json:"id"`
	SupplierCode string `json:"supplier_code"`
	SupplierName string `json:"supplier_name"`
	Nickname     string `json:"nickname"`
	Status       string `json:"status"`
}

type PaginationResponse struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type SupplierListResponse struct {
	Data       []SupplierResponse `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
