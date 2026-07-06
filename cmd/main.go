package main

import (
	"github.com/rraanggaaaa/erp_backend/pkg/config"
	"github.com/rraanggaaaa/erp_backend/pkg/routes"
)

func main() {
	config.ConnectDatabase()

	router := routes.SetupRouter()

	router.Run(":8080")
}
