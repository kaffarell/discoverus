package ports

import (
	"github.com/kaffarell/discoverus/pkg/application/core/instance"
	"github.com/kaffarell/discoverus/pkg/application/core/service"
)

// StoragePort is the port for a db adapter
type StoragePort interface {
	AddService(service service.Service) error
	GetService(serviceId string) (service.Service, error)
	GetAllServices() ([]service.Service, error)
	DeleteService(serviceId string) error

	AddInstance(serviceId string, instance instance.Instance) error
	GetInstancesOfService(serviceId string) ([]instance.Instance, error)
	GetSpecificInstance(instanceId string) (instance.Instance, error)
	GetAllInstances() ([]instance.Instance, error)
	DeleteInstance(serviceId string, instanceId string) error
}
