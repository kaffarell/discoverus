package service

import (
	"github.com/kaffarell/discoverus/pkg/ports"
	"github.com/kaffarell/discoverus/pkg/types"
)

type InstanceArray []types.Instance

func NewService(name string, serviceType string, healthCheckUrl string) bool {
	newService := types.Service{
		Id:             name,
		ServiceType:    serviceType,
		HealthCheckUrl: healthCheckUrl,
	}
	// Add service to storage
	ports.StorageInstance.AddService(newService)

	return true
}

func GetService(serviceId string) (types.Service, error) {
	service, err := ports.StorageInstance.GetService(serviceId)
	return service, err
}

func GetServices() []string {
	keys, _ := ports.StorageInstance.GetRegistry()
	return keys
}

func AddInstance(serviceName string, instance types.Instance) bool {
	err := ports.StorageInstance.AddInstance(serviceName, instance)
	if err != nil {
		return false
	}
	return true
}

func GetInstances(serviceName string) (InstanceArray, error) {
	array, err := ports.StorageInstance.GetInstances(serviceName)
	return array, err

}
