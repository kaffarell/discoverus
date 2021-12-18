package types

type Service struct {
	Id             string `json:"id"`
	ServiceType    string `json:"serviceType"`
	HealthCheckUrl string `json:"healthCheckUrl"`
}
