//go:build wireinject
// +build wireinject

package Route

import (
	"database/sql"
	"github.com/google/wire"
	"ralali-crud-cake-test/Controller"
	"ralali-crud-cake-test/Repository"
	"ralali-crud-cake-test/Services"
)

func CakeDI(db *sql.DB) *Controller.CakeControllerHandler {
	panic(wire.Build(wire.NewSet(
		Repository.CakeRepositoryControllerProvider,
		Services.CakeServicesControllerProvider,
		Controller.CakeControllerControllerProvider,

		wire.Bind(new(Controller.ICakeControllerHandler), new(*Controller.CakeControllerHandler)),
		wire.Bind(new(Services.ICakeServicesHandler), new(*Services.CakeServicesHandler)),
		wire.Bind(new(Repository.ICakeRepositoryHandler), new(*Repository.CakeRepositoryHandler)),
	),
	))
	return &Controller.CakeControllerHandler{}
}
