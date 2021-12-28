package main

import (
	"log"

	// application
	"github.com/kaffarell/discoverus/pkg/application/api"
	"github.com/kaffarell/discoverus/pkg/application/config"

	// adapters
	"github.com/kaffarell/discoverus/pkg/adapters/framework/left/rest"
	"github.com/kaffarell/discoverus/pkg/adapters/framework/right/storage"
)

func main() {

	// Initiating Database Adapter
	// Based on the DbPort interface
	dbAdapter := storage.NewAdapter()

	// Creating Configuration
	config := config.NewConfiguration(90, 10)

	// Initiating application
	// Passing the previously created dbAdapter, based on the DbPort interface
	applicationAPI := api.NewApplication(dbAdapter, config)
	log.Println("Starting application!")

	// Initiating RestApi
	// The restAPI is using the application, which is based on the ApiPort interface
	restAdapter := rest.NewAdapter(*applicationAPI)
	restAdapter.Run()
}
