package ports

import (
	"github.com/kaffarell/discoverus/pkg/types"
)

var StorageInstance Storage

type Storage interface {
	New()
	AddService(service types.Service) error
	AddInstance(serviceId string, instance types.Instance) error
	RemoveInstance(serviceId string, instanceId string) error
	GetInstances(serviceId string) ([]types.Instance, error)
	GetService(serviceId string) (types.Service, error)
	GetRegistry() ([]string, error)
}
