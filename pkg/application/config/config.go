package config

type Configuration struct {
	InstanceTimeout       int64
	InstanceTimeoutMargin int64
}

func NewConfiguration(instanceTimout int64, instanceTimoutMargin int64) Configuration {
	return Configuration{InstanceTimeout: instanceTimout, InstanceTimeoutMargin: instanceTimoutMargin}
}
