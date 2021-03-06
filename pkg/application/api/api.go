package api

import (
	"time"

	"github.com/kaffarell/discoverus/pkg/application/config"
	"github.com/kaffarell/discoverus/pkg/application/core/instance"
	"github.com/kaffarell/discoverus/pkg/application/core/service"
	"github.com/kaffarell/discoverus/pkg/ports"
	"github.com/sirupsen/logrus"
)

// Application implements the APIPort interface
type Application struct {
	logger *logrus.Logger
	db     ports.StoragePort
	config config.Configuration
}

// NewApplication creates a new Application
func NewApplication(logger *logrus.Logger, db ports.StoragePort, config config.Configuration) *Application {
	newApplication := Application{logger: logger, db: db, config: config}
	newApplication.initTicker()
	return &newApplication
}

// Initiate ticker
func (a Application) initTicker() {
	// The ticker checks if a instance has sent a heartbeat in the last 90 seconds
	// If no heartbeat has been sent, the instance will be deleted
	ticker := time.NewTicker(time.Duration(a.config.UpdateInterval) * time.Second)
	quit := make(chan struct{})
	go a.ticker(*ticker, quit)
}

// Runs every few seconds and calls another function
func (a Application) ticker(ticker time.Ticker, quit chan struct{}) {
	for {
		select {
		case <-ticker.C:
			a.logger.Info("Checking instances for inactivity...")
			a.checkInstances()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

// Created:
// 58s + 25s = 83
// Removed: 15:37:25

func (a Application) checkInstances() {
	allInstances, err := a.db.GetAllInstances()
	if err != nil {
		a.logger.Error(err)
	}

	// Get current unix time
	currentTime := time.Now().Unix()

	instancesToBeRemoved := make([]instance.Instance, 0)

	// Go through all instances and check unix time
	for _, v := range allInstances {
		if v.LastHeartbeat < (currentTime - a.config.InstanceTimeout) {
			instancesToBeRemoved = append(instancesToBeRemoved, v)
		}
	}

	// Remove all instances
	for _, v := range instancesToBeRemoved {
		// Remove instance
		a.logger.Info("Removing instance " + v.Id + " because of inactivity")
		err := a.DeleteInstance(v.ServiceId, v.Id)
		if err != nil {
			a.logger.Error("Error removing instance")
			a.logger.Error(err)
		} else {
			a.logger.Info("Removed instance: " + v.Id)
		}

	}

}

func (a Application) InsertService(service service.Service) error {
	return a.db.AddService(service)
}

func (a Application) GetService(serviceId string) (service.Service, error) {
	service, err := a.db.GetService(serviceId)
	return service, err
}

func (a Application) GetServices() ([]service.Service, error) {
	keys, err := a.db.GetAllServices()
	if err != nil {
		a.logger.Error(err)
	}
	return keys, err
}

func (a Application) AddInstance(serviceName string, instance instance.Instance) bool {
	err := a.db.AddInstance(serviceName, instance)
	a.logger.Infof("Added Instance %s to service %s", instance.Id, serviceName)
	if err != nil {
		return false
	}
	return true
}

func (a Application) GetInstancesOfService(serviceName string) ([]instance.Instance, error) {
	array, err := a.db.GetInstancesOfService(serviceName)
	if err != nil {
		a.logger.Error(err)
	}
	return array, nil
}

func (a Application) GetInstance(instanceId string) (instance.Instance, error) {
	instanceObject, err := a.db.GetSpecificInstance(instanceId)
	return instanceObject, err
}

func (a Application) GetAllInstances() ([]instance.Instance, error) {
	instances, err := a.db.GetAllInstances()
	return instances, err
}

func (a Application) DeleteInstance(serviceId string, instanceId string) error {
	err := a.db.DeleteInstance(serviceId, instanceId)
	return err
}
