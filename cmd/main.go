package main

import (
	"log"

	// application
	"github.com/kaffarell/discoverus/pkg/application/api"

	// adapters
	"github.com/kaffarell/discoverus/pkg/adapters/framework/left/rest"
	"github.com/kaffarell/discoverus/pkg/adapters/framework/right/db"
)

func main() {
	var err error

	dbAdapter, err := db.NewAdapter()
	if err != nil {
		log.Fatalf("Failed to initiate db connection: %v", err)
	}

	// NOTE: The application's right side port for driven
	// adapters, in this case, a db adapter.
	// Therefore the type for the dbAdapter parameter
	// that is to be injected into the NewApplication will
	// be of type DbPort
	applicationAPI := api.NewApplication(dbAdapter)

	// NOTE: We use dependency injection to give the grpc
	// adapter access to the application, therefore
	// the location of the port is inverted. That is
	// the grpc adapter accesses the hexagon's driving port at the
	// application boundary via dependency injection,
	// therefore the type for the applicaitonAPI parameter
	// that is to be injected into the gRPC adapter will
	// be of type APIPort which is our hexagons left side
	// port for driving adapters
	restAdapter := rest.NewAdapter(*applicationAPI)
	restAdapter.Run()
}
