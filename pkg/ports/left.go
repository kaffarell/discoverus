package ports

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	PostRegister(a, b int32)
	GetInstances(a, b int32)
	PutRenew(a, b int32)
	GetHC(a, b int32)
}
