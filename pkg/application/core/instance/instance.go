package instance

type Instance struct {
	Id            string `json:"id"`
	IP            string `json:"ip"`
	Port          int    `json:"port"`
	LastHeartbeat int64  `json:"lastHeartbeat"`
}

func NewInstance(instanceId string, ip string, port int, lastHeartbeat int64) Instance {
	return Instance{
		Id:            instanceId,
		IP:            ip,
		Port:          port,
		LastHeartbeat: lastHeartbeat,
	}
}
