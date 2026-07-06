package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rraanggaaaa/erp_backend/pkg/handler"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	// Trust localhost only (development)
	router.SetTrustedProxies(nil)

	// Handler
	supplierHandler := handler.NewSupplierHandler()

	// API Versioning
	api := router.Group("/api/v1")
	{
		// =========================
		// Health Check
		// =========================
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Backend running successfully",
			})
		})

		// =========================
		// Supplier
		// =========================
		api.GET("/suppliers", supplierHandler.GetAll)

		api.GET("/suppliers/:id", supplierHandler.GetByID)

		api.POST("/suppliers", supplierHandler.Create)

		api.PUT("/suppliers/:id", supplierHandler.Update)

		api.DELETE("/suppliers/:id", supplierHandler.Delete)
	}

	return router
}
