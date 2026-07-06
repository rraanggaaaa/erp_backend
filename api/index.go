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
	initErr error
)

func initHandler() {
	initErr = config.ConnectDatabase()
	if initErr == nil {
		handler = routes.SetupRouter()
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initHandler)

	if initErr != nil {
		http.Error(w, initErr.Error(), http.StatusInternalServerError)
		return
	}

	handler.ServeHTTP(w, r)
}
