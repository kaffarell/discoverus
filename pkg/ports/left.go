package ports

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	PostRegister(writer interface{}, req interface{}, parameter interface{})
	GetInstances(writer interface{}, req interface{}, parameter interface{})
	GetServices(writer interface{}, req interface{}, parameter interface{})
	PutRenew(writer interface{}, req interface{}, parameter interface{})
	GetHC(writer interface{}, req interface{}, parameter interface{})
}
