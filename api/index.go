package api

import (
	"net/http"

	"github.com/rraanggaaaa/erp_backend/internal/config"
	"github.com/rraanggaaaa/erp_backend/internal/routes"
)

var handler http.Handler

func init() {
	config.ConnectDatabase()
	handler = routes.SetupRouter()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	handler.ServeHTTP(w, r)
}
