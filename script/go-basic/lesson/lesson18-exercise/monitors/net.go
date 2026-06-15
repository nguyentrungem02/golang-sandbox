package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

type NetMonitor struct {
}

func (n *NetMonitor) Name() string {
	return "Network"
}

func (n *NetMonitor) Check(ctx context.Context) (string, bool) {
	netStat, err := net.IOCountersWithContext(ctx, false)
	if err != nil && len(netStat) == 0 {
		return fmt.Sprintf("[Network Monitor] Could not retrieve Network info: %v\n", err), false
	}

	value := fmt.Sprintf("Send: %d KB, Rev: %d KB", netStat[0].BytesSent/1024, netStat[0].BytesRecv/1024)
	return value, false
}
