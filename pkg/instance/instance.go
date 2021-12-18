package instance

import "github.com/kaffarell/discoverus/pkg/types"

func NewInstance(instanceId string, ip string, port int, lastHearbeat int64) types.Instance {
	return types.Instance{
		Id:            instanceId,
		IP:            ip,
		Port:          port,
		LastHeartbeat: lastHearbeat,
	}
}
