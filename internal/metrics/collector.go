package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// Collector retrieves system metrics.
type Collector struct{}

// NewCollector creates a new Collector instance.
func NewCollector() Collector {
	return Collector{}
}

// Snapshot takes a snapshot of the current CPU and RAM metrics.
func (c Collector) Snapshot() (Snapshot, error) {
	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil {
		return Snapshot{}, err
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return Snapshot{}, err
	}

	snapshot := Snapshot{
		Timestamp:   time.Now(),
		CPUUsage:    cpuPercentages[0],
		MemoryUsage: vmStat.UsedPercent,
		DiskUsage:   0, // TODO: Implement disk usage
	}

	return snapshot, nil
}
