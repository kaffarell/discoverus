package api

import (
	"github.com/kaffarell/discoverus/pkg/application/core/instance"
	"github.com/kaffarell/discoverus/pkg/application/core/service"
	"github.com/kaffarell/discoverus/pkg/ports"
)

// Application implements the APIPort interface
type Application struct {
	db ports.DbPort
}

// NewApplication creates a new Application
func NewApplication(db ports.DbPort) *Application {
	return &Application{db: db}
}

func (a Application) InsertService(service service.Service) error {
	return a.db.AddService(service)
}

func (a Application) GetService(serviceId string) (service.Service, error) {
	service, err := a.db.GetService(serviceId)
	return service, err
}

func (a Application) GetServices() []string {
	keys, _ := a.db.GetRegistry()
	return keys
}

func (a Application) AddInstance(serviceName string, instance instance.Instance) bool {
	err := a.db.AddInstance(serviceName, instance)
	if err != nil {
		return false
	}
	return true
}

func (a Application) GetInstances(serviceName string) ([]instance.Instance, error) {
	array, err := a.db.GetInstances(serviceName)
	return array, err

}
