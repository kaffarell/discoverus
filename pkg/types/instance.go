package types

type Instance struct {
	Id            string `json:"id"`
	IP            string `json:"ip"`
	Port          int    `json:"port"`
	LastHeartbeat int64  `json:"lastHeartbeat"`
}
