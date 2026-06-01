package metricscomponent

import (
	"errors"
	"testing"
	"time"

	"axiomtui/internal/metrics"
)

func TestUpdateSnapshotLoaded(t *testing.T) {
	snapshot := metrics.Snapshot{Timestamp: time.Now(), CPUUsage: 12.5, MemoryUsage: 48.25}

	updated, cmd := Model{}.Update(SnapshotLoaded{Snapshot: snapshot})

	if cmd != nil {
		t.Fatalf("cmd = %v, want nil", cmd)
	}
	if !updated.Loaded {
		t.Fatal("Loaded = false, want true")
	}
	if updated.Error != nil {
		t.Fatalf("Error = %v, want nil", updated.Error)
	}
	if updated.Snapshot.CPUUsage != 12.5 || updated.Snapshot.MemoryUsage != 48.25 {
		t.Fatalf("Snapshot = %#v, want CPU/RAM values", updated.Snapshot)
	}
}

func TestUpdateSnapshotError(t *testing.T) {
	err := errors.New("metrics failed")

	updated, cmd := Model{}.Update(SnapshotError{Err: err})

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
