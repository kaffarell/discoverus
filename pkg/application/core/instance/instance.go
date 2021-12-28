package instance

type Instance struct {
	Id            string `json:"id"`
	ServiceId     string `json:"serviceid"`
	IP            string `json:"ip"`
	Port          int    `json:"port"`
	LastHeartbeat int64  `json:"lastHeartbeat"`
}

func NewInstance(instanceId string, serviceId string, ip string, port int, lastHeartbeat int64) Instance {
	return Instance{
		Id:            instanceId,
		ServiceId:     serviceId,
		IP:            ip,
		Port:          port,
		LastHeartbeat: lastHeartbeat,
	}
}
