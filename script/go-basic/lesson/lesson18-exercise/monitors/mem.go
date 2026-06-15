package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemMonitor struct {
}

func (cm *MemMonitor) Name() string {
	return "Memory"
}

func (cm *MemMonitor) Check(ctx context.Context) (string, bool) {
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Memory Monitor] Could not retrieve Memory info: %v\n", err), false
	}

	value := fmt.Sprintf("%.2f%%", vmStat.UsedPercent)

	return value, vmStat.UsedPercent > 60
}
