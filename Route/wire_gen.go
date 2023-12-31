// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package Route

import (
	"database/sql"
	"ralali-crud-cake-test/Controller"
	"ralali-crud-cake-test/Repository"
	"ralali-crud-cake-test/Services"
)

// Injectors from Wire.go:

func CakeDI(db *sql.DB) *Controller.CakeControllerHandler {
	cakeRepositoryHandler := Repository.CakeRepositoryControllerProvider(db)
	cakeServicesHandler := Services.CakeServicesControllerProvider(cakeRepositoryHandler)
	cakeControllerHandler := Controller.CakeControllerControllerProvider(cakeServicesHandler)
	return cakeControllerHandler
}
