package rest

import (
	"github.com/kaffarell/discoverus/pkg/application/api"
)

// Adapter implements the GRPCPort interface
type Adapter struct {
	application api.Application
}

// NewAdapter creates a new Adapter
func NewAdapter(application api.Application) *Adapter {
	return &Adapter{application: application}
}

func (a Adapter) Run() {
}

func (a Adapter) PostRegister() {

}
func (a Adapter) GetInstances() {

}
func (a Adapter) PutRenew() {

}
func (a Adapter) GetHC() {

}
