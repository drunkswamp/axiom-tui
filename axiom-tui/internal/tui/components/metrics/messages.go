package metricscomponent

import "axiomtui/internal/metrics"

// SnapshotLoaded is emitted when metrics are loaded successfully.
type SnapshotLoaded struct {
	Snapshot metrics.Snapshot
}

// SnapshotError is emitted when metrics loading fails.
type SnapshotError struct {
	Err error
}
