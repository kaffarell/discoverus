package ports

import (
	"github.com/kaffarell/discoverus/pkg/application/core/instance"
	"github.com/kaffarell/discoverus/pkg/application/core/service"
)

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	GetInstances(serviceId string) ([]instance.Instance, error)
	GetInstance(instanceId string) (instance.Instance, error)
	AddInstance(serviceId string, instance instance.Instance) bool
	GetServices() string
	GetService(serviceId string) (service.Service, error)
	InsertService(service service.Service) error
	DeleteInstance(serviceId string, instanceId string) error
}
