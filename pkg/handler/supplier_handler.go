package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rraanggaaaa/erp_backend/pkg/dto"
	"github.com/rraanggaaaa/erp_backend/pkg/model"
	"github.com/rraanggaaaa/erp_backend/pkg/repository"
	"github.com/rraanggaaaa/erp_backend/pkg/utils"
)

type SupplierHandler struct {
	repository *repository.SupplierRepository
}

func NewSupplierHandler() *SupplierHandler {
	return &SupplierHandler{
		repository: repository.NewSupplierRepository(),
	}
}

// ======================================
// GET ALL SUPPLIERS
// ======================================

func (h *SupplierHandler) GetAll(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	search := c.DefaultQuery("search", "")

	suppliers, total, err := h.repository.GetAll(search, page, limit)

	if err != nil {

		utils.Error(
			c,
			http.StatusInternalServerError,
			"Failed to get suppliers",
			err.Error(),
		)

		return
	}

	response := make([]dto.SupplierResponse, 0)

	for _, supplier := range suppliers {

		response = append(response, dto.SupplierResponse{
			ID:           supplier.ID.String(),
			SupplierCode: supplier.SupplierCode,
			SupplierName: supplier.SupplierName,
			Nickname:     supplier.Nickname,
			Status:       supplier.Status,
		})

	}

	utils.Success(
		c,
		http.StatusOK,
		"Success",
		gin.H{
			"items": response,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
			},
		},
	)

}

// ======================================
// GET DETAIL SUPPLIER
// ======================================

func (h *SupplierHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	supplier, err := h.repository.GetByID(id)

	if err != nil {

		utils.Error(
			c,
			http.StatusNotFound,
			"Supplier not found",
			nil,
		)

		return

	}

	response := dto.SupplierResponse{

		ID:           supplier.ID.String(),
		SupplierCode: supplier.SupplierCode,
		SupplierName: supplier.SupplierName,
		Nickname:     supplier.Nickname,
		Status:       supplier.Status,
	}

	utils.Success(
		c,
		http.StatusOK,
		"Success",
		response,
	)

}

// ======================================
// CREATE SUPPLIER
// ======================================

func (h *SupplierHandler) Create(c *gin.Context) {

	var request dto.CreateSupplierRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		utils.Error(
			c,
			http.StatusBadRequest,
			"Validation Error",
			err.Error(),
		)

		return
	}

	lastCode, _ := h.repository.GetLastSupplierCode()

	lastNumber := 0

	if len(lastCode) >= 8 {

		lastNumber, _ = strconv.Atoi(lastCode[2:])

	}

	supplier := model.Supplier{

		SupplierCode: utils.GenerateSupplierCode(lastNumber),

		SupplierName: request.SupplierName,

		Nickname: request.Nickname,

		Status: request.Status,
	}

	id, err := h.repository.Create(supplier)

	if err != nil {

		utils.Error(
			c,
			http.StatusInternalServerError,
			"Failed to create supplier",
			err.Error(),
		)

		return

	}

	utils.Success(
		c,
		http.StatusCreated,
		"Supplier created successfully",
		gin.H{
			"id":            id,
			"supplier_code": supplier.SupplierCode,
		},
	)

}

func (h *SupplierHandler) Update(c *gin.Context) {

	id := c.Param("id")

	var request dto.CreateSupplierRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		utils.Error(
			c,
			http.StatusBadRequest,
			"Validation Error",
			err.Error(),
		)

		return
	}

	supplier := model.Supplier{
		SupplierName: request.SupplierName,
		Nickname:     request.Nickname,
		Status:       request.Status,
	}

	err := h.repository.Update(id, supplier)

	if err != nil {

		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
			nil,
		)

		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Supplier updated successfully",
		nil,
	)
}

func (h *SupplierHandler) Delete(c *gin.Context) {

	id := c.Param("id")

	err := h.repository.Delete(id)

	if err != nil {

		utils.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
			nil,
		)

		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Supplier deleted successfully",
		nil,
	)

}
