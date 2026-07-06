package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/rraanggaaaa/erp_backend/pkg/config"
	"github.com/rraanggaaaa/erp_backend/pkg/model"
)

type SupplierRepository struct{}

func NewSupplierRepository() *SupplierRepository {
	return &SupplierRepository{}
}

// ======================================================
// GET ALL SUPPLIERS
// Support:
// - Search
// - Pagination
// ======================================================

func (r *SupplierRepository) GetAll(
	search string,
	page int,
	limit int,
) ([]model.Supplier, int, error) {

	offset := (page - 1) * limit

	query := `
	SELECT
		id,
		supplier_code,
		supplier_name,
		nickname,
		logo_url,
		status,
		created_at,
		updated_at
	FROM suppliers
	WHERE supplier_name ILIKE '%' || $1 || '%'
	ORDER BY supplier_name ASC
	LIMIT $2
	OFFSET $3
	`

	rows, err := config.DB.Query(
		context.Background(),
		query,
		search,
		limit,
		offset,
	)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	suppliers := make([]model.Supplier, 0)

	for rows.Next() {

		var supplier model.Supplier

		err := rows.Scan(
			&supplier.ID,
			&supplier.SupplierCode,
			&supplier.SupplierName,
			&supplier.Nickname,
			&supplier.LogoURL,
			&supplier.Status,
			&supplier.CreatedAt,
			&supplier.UpdatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		suppliers = append(suppliers, supplier)

	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	var total int

	countQuery := `
	SELECT COUNT(*)
	FROM suppliers
	WHERE supplier_name ILIKE '%' || $1 || '%'
	`

	err = config.DB.QueryRow(
		context.Background(),
		countQuery,
		search,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	return suppliers, total, nil
}

// ======================================================
// GET SUPPLIER BY ID
// ======================================================

func (r *SupplierRepository) GetByID(id string) (*model.Supplier, error) {

	query := `
	SELECT
		id,
		supplier_code,
		supplier_name,
		nickname,
		logo_url,
		status,
		created_at,
		updated_at
	FROM suppliers
	WHERE id = $1
	`

	var supplier model.Supplier
	var uuidID uuid.UUID

	err := config.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&uuidID,
		&supplier.SupplierCode,
		&supplier.SupplierName,
		&supplier.Nickname,
		&supplier.LogoURL,
		&supplier.Status,
		&supplier.CreatedAt,
		&supplier.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("supplier not found")
	}

	supplier.ID = uuidID

	return &supplier, nil
}

// ======================================================
// GET LAST SUPPLIER CODE
// ======================================================

func (r *SupplierRepository) GetLastSupplierCode() (string, error) {

	query := `
	SELECT supplier_code
	FROM suppliers
	ORDER BY supplier_code DESC
	LIMIT 1
	`

	var supplierCode string

	err := config.DB.QueryRow(
		context.Background(),
		query,
	).Scan(&supplierCode)

	// Jika tabel masih kosong
	if err != nil {
		return "61000000", nil
	}

	return supplierCode, nil
}

// ======================================================
// CREATE SUPPLIER
// Database akan membuat UUID otomatis
// ======================================================

func (r *SupplierRepository) Create(
	supplier model.Supplier,
) (string, error) {

	query := `
	INSERT INTO suppliers
	(
		supplier_code,
		supplier_name,
		nickname,
		logo_url,
		status
	)
	VALUES
	($1,$2,$3,$4,$5)
	RETURNING id
	`

	var id string

	err := config.DB.QueryRow(
		context.Background(),
		query,
		supplier.SupplierCode,
		supplier.SupplierName,
		supplier.Nickname,
		supplier.LogoURL,
		supplier.Status,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// ======================================================
// UPDATE SUPPLIER
// ======================================================

func (r *SupplierRepository) Update(
	id string,
	supplier model.Supplier,
) error {

	query := `
	UPDATE suppliers
	SET
		supplier_name = $1,
		nickname = $2,
		logo_url = $3,
		status = $4,
		updated_at = NOW()
	WHERE id = $5
	`

	commandTag, err := config.DB.Exec(
		context.Background(),
		query,
		supplier.SupplierName,
		supplier.Nickname,
		supplier.LogoURL,
		supplier.Status,
		id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("supplier not found")
	}

	return nil
}

// ======================================================
// DELETE SUPPLIER
// ======================================================

func (r *SupplierRepository) Delete(id string) error {

	query := `
	DELETE FROM suppliers
	WHERE id = $1
	`

	commandTag, err := config.DB.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("supplier not found")
	}

	return nil
}
