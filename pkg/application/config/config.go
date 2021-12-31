package config

type Configuration struct {
	// When an instance doesn't send any heartbeat after the
	// InstanceTimeout it gets removed
	InstanceTimeout int64
	// The application checks every UpdateInterval seconds if any
	// instance needs to be removed
	UpdateInterval int64
}

func NewConfiguration(instanceTimout int64, updateInterval int64) Configuration {
	return Configuration{
		InstanceTimeout: instanceTimout,
		UpdateInterval:  updateInterval,
	}
}
