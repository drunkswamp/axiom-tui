package dockercomponent

import (
	"errors"
	"testing"

	"axiomtui/internal/domain"
)

func TestUpdateServicesLoadedMarksModelLoaded(t *testing.T) {
	model := Model{}
	services := []domain.Service{{ID: "abc123", Name: "web", Type: domain.ServiceTypeDocker}}

	updated, cmd := model.Update(ServicesLoaded{Services: services})

	if cmd != nil {
		t.Fatalf("cmd = %v, want nil", cmd)
	}
	if !updated.Loaded {
		t.Fatal("Loaded = false, want true")
	}
	if updated.Error != nil {
		t.Fatalf("Error = %v, want nil", updated.Error)
	}
	if len(updated.Services) != 1 || updated.Services[0].Name != "web" {
		t.Fatalf("Services = %#v, want web service", updated.Services)
	}
}

func TestUpdateServicesErrorMarksModelLoaded(t *testing.T) {
	model := Model{Services: []domain.Service{{ID: "old", Name: "old"}}}
	err := errors.New("docker unavailable")

	updated, cmd := model.Update(ServicesError{Err: err})

	if cmd != nil {
		t.Fatalf("cmd = %v, want nil", cmd)
	}
	if !updated.Loaded {
		t.Fatal("Loaded = false, want true")
	}
	if updated.Error != err {
		t.Fatalf("Error = %v, want %v", updated.Error, err)
	}
}
