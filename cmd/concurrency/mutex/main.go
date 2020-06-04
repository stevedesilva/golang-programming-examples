package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	funcWithMutex()
}

func funcWithMutex() {
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	fmt.Printf(">>>Goroutines: %d\n", runtime.NumGoroutine())
	var (
		counter int
		wg      sync.WaitGroup
		numgr   int = 100
		mu      sync.Mutex
	)

	wg.Add(numgr)
	for i := 0; i < numgr; i++ {
		go func(i int) {
			mu.Lock()
			fmt.Printf("gr %d\n", i)
			v := counter
			v++
			counter = v
			fmt.Printf("gr %d\tcounter %v\n", i, v)
			mu.Unlock()
			wg.Done()

		}(i)
	}
	fmt.Printf("<<<Goroutines: %d\n", runtime.NumGoroutine())
	wg.Wait()

	fmt.Printf("Counter %d\n", counter)
}
