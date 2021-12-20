package service

import "github.com/kaffarell/discoverus/pkg/application/core/instance"

type Service struct {
	Id             string `json:"id"`
	ServiceType    string `json:"serviceType"`
	HealthCheckUrl string `json:"healthCheckUrl"`
}

type InstanceArray []instance.Instance

func NewService(name string, serviceType string, healthCheckUrl string) Service {
	newService := Service{
		Id:             name,
		ServiceType:    serviceType,
		HealthCheckUrl: healthCheckUrl,
	}
	return newService
}
