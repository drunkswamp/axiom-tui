package metrics

import "time"

// Snapshot is a point-in-time snapshot of key node metrics.
type Snapshot struct {
	Timestamp   time.Time
	CPUUsage    float64
	MemoryUsage float64
	DiskUsage   float64
}
