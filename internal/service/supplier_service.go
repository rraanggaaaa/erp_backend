package service

import (
	"github.com/rraanggaaaa/erp_backend/internal/model"
	"github.com/rraanggaaaa/erp_backend/internal/repository"
)

type SupplierService struct {
	repository *repository.SupplierRepository
}

func NewSupplierService() *SupplierService {
	return &SupplierService{
		repository: repository.NewSupplierRepository(),
	}
}

// ===============================
// GET ALL SUPPLIERS
// ===============================

func (s *SupplierService) GetAll(search string, page int, limit int) ([]model.Supplier, int, error) {
	return s.repository.GetAll(search, page, limit)
}

// ===============================
// GET SUPPLIER BY ID
// ===============================

func (s *SupplierService) GetByID(id string) (*model.Supplier, error) {
	return s.repository.GetByID(id)
}

// ===============================
// GET LAST SUPPLIER CODE
// ===============================

func (s *SupplierService) GetLastSupplierCode() (string, error) {
	return s.repository.GetLastSupplierCode()
}

// ===============================
// CREATE SUPPLIER
// ===============================

func (s *SupplierService) Create(supplier model.Supplier) (string, error) {
	return s.repository.Create(supplier)
}

// ===============================
// UPDATE SUPPLIER
// ===============================

func (s *SupplierService) Update(id string, supplier model.Supplier) error {
	return s.repository.Update(id, supplier)
}

// ===============================
// DELETE SUPPLIER
// ===============================

func (s *SupplierService) Delete(id string) error {
	return s.repository.Delete(id)
}
