package processor

import (
	"context"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
	"trungem.com/hoc-golang/models"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStat, m models.Monitor) {
	defer wg.Done()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			value, alert := m.Check(ctx)

			stat := models.SystemStat{
				Name:  m.Name(),
				Value: value,
			}

			statCh <- stat

			if alert {
				LogAlert(stat)
			}
		}
	}
}

func GetTopProcesses(ctx context.Context) string {
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve mem list: %v\n", err)
	}

	totalMemory := vmStat.Total

	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve process list: %v\n", err)
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	var cpuList, memList []models.ProcStat
	procCh := make(chan models.ProcStat, len(processes))

	for _, p := range processes {
		wg.Add(1)
		go func(proc *process.Process) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				name, err := p.NameWithContext(ctx)
				if err != nil {
					return
				}
				cpuPercent, err := p.CPUPercentWithContext(ctx)
				if err != nil {
					return
				}

				// RSS: Resident Set Size => Ram thực sự mà tiến trình đang sử dụng
				// VMS: Virtual Memory Size => Tổng bộ nhớ ảo mà hệ điều hành cấp phát cho tiến trình
				memInfo, err := p.MemoryInfoWithContext(ctx)
				if err != nil {
					return
				}

				ramPercent := float64(memInfo.RSS) / float64(totalMemory) * 100

				// createTime đang có dữ liệu là milliseconds
				createTime, err := proc.CreateTimeWithContext(ctx)
				if err != nil {
					return
				}

				// Time Unix là tổng số giây từ mốc thời gian 1970-01-01 00:00:00 UTC đến hiện tại
				// Vì createTime là milliseconds nên cần phải chia 1000 để convert thành second
				runningTime := time.Since(time.Unix(createTime/1000, 0))

				if cpuPercent > 1 || ramPercent > 1 {
					mu.Lock()
					procStat := models.ProcStat{
						PID:         proc.Pid,
						Name:        name,
						CPU:         cpuPercent,
						Memory:      memInfo.RSS,
						RamPercent:  ramPercent,
						RunningTime: runningTime,
					}

					procCh <- procStat
					mu.Unlock()
				}
			}

		}(p)
	}

	go func() {
		wg.Wait()
		close(procCh)
	}()

	for stat := range procCh {
		if stat.CPU > 1 {
			cpuList = append(cpuList, stat)
		}

		if stat.Memory > 1 {
			memList = append(memList, stat)
		}
	}

	sort.Slice(cpuList, func(i, j int) bool {
		return cpuList[i].CPU > cpuList[j].CPU
	})

	sort.Slice(memList, func(i, j int) bool {
		return memList[i].RamPercent > memList[j].RamPercent
	})

	output := "== Top 5 CPU consuming processes == \n"
	for i := 0; i < len(cpuList) && i < 5; i++ {
		output += fmt.Sprintf("%d. [%d] %s - CPU: %.2f%% - RAM: %.2f MB (%.2f%%) - Running: %s \n", i+1, cpuList[i].PID, cpuList[i].Name, cpuList[i].CPU, float64(cpuList[i].Memory)/1024.0/1024.0, cpuList[i].RamPercent, cpuList[i].RunningTime)
	}

	output += "== Top 5 RAM consuming processes == \n"
	for i := 0; i < len(memList) && i < 5; i++ {
		output += fmt.Sprintf("%d. [%d] %s - CPU: %.2f%% - RAM: %.2f MB (%.2f%%) - Running: %s \n", i+1, memList[i].PID, memList[i].Name, memList[i].CPU, float64(memList[i].Memory)/1024.0/1024.0, memList[i].RamPercent, cpuList[i].RunningTime)
	}

	ExportToCSV(cpuList, memList)

	return output
}

func ExportToCSV(cpuList, memList []models.ProcStat) {
	file, err := os.OpenFile("process_stats.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[Export To CSV] Could not open file: %v\n", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	if stat, err := file.Stat(); err == nil && stat.Size() == 0 {
		file.WriteString("Timestamp,PID,Name,CPU (%),RAM (MB),RAM (%),Running Time \n")
	}

	timestamp := time.Now().Format(time.RFC3339)
	for i := 0; i < len(cpuList) && i < 5; i++ {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s \n", timestamp, cpuList[i].PID, cpuList[i].Name, cpuList[i].CPU, float64(cpuList[i].Memory)/1024.0/1024.0, cpuList[i].RamPercent, cpuList[i].RunningTime)
		file.WriteString(line)
	}

	for i := 0; i < len(memList) && i < 5; i++ {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s \n", timestamp, memList[i].PID, memList[i].Name, memList[i].CPU, float64(memList[i].Memory)/1024.0/1024.0, memList[i].RamPercent, memList[i].RunningTime)
		file.WriteString(line)
	}
}

func LogAlert(stat models.SystemStat) {
	models.StatMutex.Lock()
	defer models.StatMutex.Unlock()

	file, err := os.OpenFile("alert.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[Log Alert] Failed to write log: %v\n", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	timestamp := time.Now().Format(time.RFC3339)
	logLine := fmt.Sprintf("[%s] ALERT: %s = %s \n", timestamp, stat.Name, stat.Value)
	file.WriteString(logLine)
}
