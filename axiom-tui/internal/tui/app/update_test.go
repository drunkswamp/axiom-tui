package app

import "testing"

func TestInitReturnsInitialCommand(t *testing.T) {
	model, err := NewAppModel()
	if err != nil {
		t.Fatalf("NewAppModel() error = %v", err)
	}

	if cmd := model.Init(); cmd == nil {
		t.Fatal("Init() returned nil command, want initial load command")
	}
}
