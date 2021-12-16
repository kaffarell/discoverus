package service

import (
	"errors"
)

type InstanceArray []Instance

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

func AddInstance(serviceName string, instance Instance) bool {
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

func NewInstance(instanceId int, ip string, port int) Instance {
	return Instance{
		InstanceId: instanceId,
		IP:         ip,
		Port:       port,
	}
}

type Instance struct {
	InstanceId int    `json:"instanceId"`
	IP         string `json:"ip"`
	Port       int    `json:"port"`
}
