package service

import (
	"errors"
    "github.com/kaffarell/discoverus/internal/instance"
)

type InstanceArray []instance.Instance

var Services map[Service]InstanceArray

type Service struct {
	Name           string
	ServiceType    string
	HealthCheckUrl string
}


func NewService(name string, serviceType string, healthCheckUrl string) bool {
	newService := Service{
		Name:           name,
		ServiceType:    serviceType,
		HealthCheckUrl: healthCheckUrl,
	}
	if Services[newService] == nil {
		Services[newService] = make(InstanceArray, 0)
	} else {
		return false
	}

	return true
}

func GetServices() []Service {
	keys := make([]Service, len(Services))

	i := 0
	for k := range Services {
		keys[i] = k
		i++
	}
	return keys
}

/*
	Searches for service in map and returns the struct
*/
func GetService(serviceName string) (*Service, error) {
	for k := range Services {
		if k.Name == serviceName {
			return &k, nil
		}

	}
	return nil, errors.New("Service not found")
}

func AddInstance(serviceName string, instance instance.Instance) bool {
	search, error := GetService(serviceName)
	if error != nil {
		return false
	}
	Services[*search] = append(Services[*search], instance)

	return true
}

func GetInstances(serviceName string) (InstanceArray, error) {
	search, error := GetService(serviceName)
	if error != nil {
		return nil, errors.New("No Service found")
	}
	return Services[*search], nil

}


