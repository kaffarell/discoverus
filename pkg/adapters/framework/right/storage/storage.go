package storage

import (
	"errors"
	"sync"

	// FIXME: Maybe use interface{} as type and use Services in application/api
	"github.com/kaffarell/discoverus/pkg/application/core/instance"
	"github.com/kaffarell/discoverus/pkg/application/core/service"
)

// Adapter implements the DbPort interface
type Adapter struct {
	allInstancesMutex sync.Mutex
	allInstances      []instance.Instance

	allServicesMutex sync.Mutex
	allServices      []service.Service
}

// NewAdapter creates a new Adapter
func NewAdapter() *Adapter {
	return &Adapter{
		allInstances: make([]instance.Instance, 0),
		allServices:  make([]service.Service, 0),
	}
}

func (a *Adapter) AddService(service service.Service) error {
	a.allServicesMutex.Lock()
	defer a.allServicesMutex.Unlock()

	a.allServices = append(a.allServices, service)
	return nil
}

func (a *Adapter) AddInstance(serviceId string, instance instance.Instance) error {
	a.allInstancesMutex.Lock()
	defer a.allInstancesMutex.Unlock()

	a.allInstances = append(a.allInstances, instance)

	return nil
}

func (a *Adapter) DeleteInstance(serviceId string, instanceId string) error {
	a.allInstancesMutex.Lock()
	defer a.allInstancesMutex.Unlock()

	for i, v := range a.allInstances {
		if v.Id == instanceId {
			// Remove this instance
			a.allInstances = append(a.allInstances[:i], a.allInstances[i+1:]...)
			// Only remove one instance
			return nil
		}
	}
	return errors.New("no instance to delete found")
}

func (a *Adapter) DeleteService(serviceId string) error {
	a.allServicesMutex.Lock()
	defer a.allServicesMutex.Unlock()

	for i, v := range a.allServices {
		if v.Id == serviceId {
			// Remove this instance
			a.allServices = append(a.allServices[:i], a.allServices[i+1:]...)
			// Only remove one instance
			return nil
		}
	}
	return errors.New("no instance to delete found")
}

func (a *Adapter) GetInstancesOfService(serviceId string) ([]instance.Instance, error) {
	a.allInstancesMutex.Lock()
	defer a.allInstancesMutex.Unlock()

	instances := make([]instance.Instance, 0)
	for _, v := range a.allInstances {
		if v.ServiceId == serviceId {
			instances = append(instances, v)
		}
	}
	return instances, nil

}

func (a *Adapter) GetSpecificInstance(instanceId string) (instance.Instance, error) {
	a.allInstancesMutex.Lock()
	defer a.allInstancesMutex.Unlock()

	for _, v := range a.allInstances {
		if v.Id == instanceId {
			return v, nil
		}
	}
	return instance.Instance{}, errors.New("no instance with id " + instanceId + " found")
}

func (a *Adapter) GetAllServices() ([]service.Service, error) {
	a.allServicesMutex.Lock()
	defer a.allServicesMutex.Unlock()

	services := make([]service.Service, len(a.allServices))
	copy(services, a.allServices)

	return services, nil
}

func (a *Adapter) GetService(serviceId string) (service.Service, error) {
	a.allServicesMutex.Lock()
	defer a.allServicesMutex.Unlock()

	for _, v := range a.allServices {
		if v.Id == serviceId {
			return v, nil
		}
	}
	return service.Service{}, errors.New("no service with id " + serviceId + " found")
}

func (a *Adapter) GetAllInstances() ([]instance.Instance, error) {
	a.allInstancesMutex.Lock()
	defer a.allInstancesMutex.Unlock()

	instances := make([]instance.Instance, len(a.allInstances))
	copy(instances, a.allInstances)

	return instances, nil
}
