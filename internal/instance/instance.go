package instance

type Instance struct {
	InstanceId int    `json:"instanceId"`
	IP         string `json:"ip"`
	Port       int    `json:"port"`
}

func NewInstance(instanceId int, ip string, port int) Instance {
	return Instance{
		InstanceId: instanceId,
		IP:         ip,
		Port:       port,
	}
}

