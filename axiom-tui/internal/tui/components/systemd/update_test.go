package systemdcomponent

import (
	"errors"
	"testing"

	"axiomtui/internal/domain"
)

func TestUpdateServicesLoaded(t *testing.T) {
	services := []domain.Service{{ID: "nginx.service", Name: "Nginx", Type: domain.ServiceTypeSystemd}}

	updated, cmd := Model{}.Update(ServicesLoaded{Services: services})

	if cmd != nil {
		t.Fatalf("cmd = %v, want nil", cmd)
	}
	if !updated.Loaded {
		t.Fatal("Loaded = false, want true")
	}
	if updated.Error != nil {
		t.Fatalf("Error = %v, want nil", updated.Error)
	}
	if len(updated.Services) != 1 || updated.Services[0].ID != "nginx.service" {
		t.Fatalf("Services = %#v, want nginx.service", updated.Services)
	}
}

func TestUpdateServicesError(t *testing.T) {
	err := errors.New("systemd failed")

	updated, cmd := Model{}.Update(ServicesError{Err: err})

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
