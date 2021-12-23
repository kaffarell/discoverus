package ports

import (
	"github.com/kaffarell/discoverus/pkg/application/core/instance"
	"github.com/kaffarell/discoverus/pkg/application/core/service"
)

// DbPort is the port for a db adapter
type DbPort interface {
	AddService(service service.Service) error
	AddInstance(serviceId string, instance instance.Instance) error
	RemoveInstance(serviceId string, instanceId string) error
	GetInstances(serviceId string) ([]instance.Instance, error)
	GetInstance(instanceId string) (instance.Instance, error)
	GetService(serviceId string) (service.Service, error)
	GetRegistry() ([]string, error)
}
