package instance

type Instance struct {
	InstanceId    string `json:"instanceId"`
	IP            string `json:"ip"`
	Port          int    `json:"port"`
	LastHeartbeat int64  `json:"lastHearbeat"`
}

func NewInstance(instanceId string, ip string, port int, lastHearbeat int64) Instance {
	return Instance{
		InstanceId:    instanceId,
		IP:            ip,
		Port:          port,
		LastHeartbeat: lastHearbeat,
	}
}
