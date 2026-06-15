package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"trungem.com/hoc-golang/models"
	"trungem.com/hoc-golang/monitors"
	"trungem.com/hoc-golang/processor"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitorList := []models.Monitor{
		&monitors.CPUMonitor{},
		&monitors.MemMonitor{},
		&monitors.NetMonitor{},
		&monitors.DiskMonitor{},
	}

	var wg sync.WaitGroup
	statCh := make(chan models.SystemStat)

	for _, monitor := range monitorList {
		wg.Add(1)
		go processor.RunMonitor(ctx, &wg, statCh, monitor)
	}

	go func() {
		for stat := range statCh {
			models.StatMutex.Lock()
			models.Stats[stat.Name] = stat
			models.StatMutex.Unlock()
		}
	}()

	printTicker := time.NewTicker(5 * time.Second)
	go func() {
		for range printTicker.C {
			fmt.Println("=== System Status ===")

			for _, stat := range models.Stats {
				models.StatMutex.Lock()
				fmt.Printf("[%s] %s\n", stat.Name, stat.Value)
				models.StatMutex.Unlock()
			}

			fmt.Println(processor.GetTopProcesses(ctx))
		}
	}()

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
	close(statCh)
	printTicker.Stop()
}
