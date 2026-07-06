package api

import (
	"net/http"
	"sync"

	"github.com/rraanggaaaa/erp_backend/pkg/config"
	"github.com/rraanggaaaa/erp_backend/pkg/routes"
)

var (
	handler http.Handler
	once    sync.Once
)

func initHandler() {
	config.ConnectDatabase()
	handler = routes.SetupRouter()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initHandler)
	handler.ServeHTTP(w, r)
}
