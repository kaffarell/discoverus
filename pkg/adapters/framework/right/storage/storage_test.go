package storage

import (
	"testing"

	"github.com/kaffarell/discoverus/pkg/application/core/service"
)

func TestAddService(t *testing.T) {
	a := NewAdapter()
	sample_service := service.NewService("test", "service", "healthurl")
	a.AddService(sample_service)

	got, err := a.GetService("test")
	if err != nil {
		t.Error("Got error")
	}
	if sample_service != got {
		t.Error("Service not inserted correctly")
	}
}
