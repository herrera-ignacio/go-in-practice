package main

import (
	"log"
	"runtime"
	"time"
)

func monitorRuntime() {
	// When monitoring starts, report the number of processors available
	log.Println("Number of CPUs:", runtime.NumCPU())
	m := &runtime.MemStats{}
	// Loop continuously, pausing for 10 seconds between each iteration
	for {
		// Logs the number of goroutines and amount of allocated memory
		r := runtime.NumGoroutine()
		log.Println("Number of goroutines", r)

		runtime.ReadMemStats(m)
		log.Println("Allocated memory", m.Alloc)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	// When the application starts, begins monitoring the application
	go monitorRuntime()

	// Creates example goroutines and memory usage while the application runs for 40 seconds
	i := 0
	for i < 40 {
		go func() {
			time.Sleep(15 * time.Second)
		}()

		i = i + 1
		time.Sleep(1 * time.Second)
	}
}
