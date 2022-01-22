package main

import (

	// application
	"github.com/kaffarell/discoverus/pkg/application/api"
	"github.com/kaffarell/discoverus/pkg/application/config"

	"net"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	// adapters
	"github.com/kaffarell/discoverus/pkg/adapters/framework/left/rest"
	"github.com/kaffarell/discoverus/pkg/adapters/framework/right/storage"
	"github.com/sirupsen/logrus"
)

func main() {

	// Initiating Database Adapter
	// Based on the DbPort interface
	dbAdapter := storage.NewAdapter()

	// Creating Configuration
	config := config.NewConfiguration(90, 15)

	// Creating logger instance
	log := logrus.New()
	conn, err := net.Dial("tcp", "logstash:5000")
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "discoverus"}))

	log.Hooks.Add(hook)
	ctx := log.WithFields(logrus.Fields{
		"method": "main",
	})
	ctx.Info("Instantiated logger")

	// Initiating application
	// Passing the previously created dbAdapter, based on the DbPort interface
	applicationAPI := api.NewApplication(log, dbAdapter, config)
	log.Info("Starting application!")

	// Initiating RestApi
	// The restAPI is using the application, which is based on the ApiPort interface
	restAdapter := rest.NewAdapter(*applicationAPI)
	restAdapter.Run()
}
