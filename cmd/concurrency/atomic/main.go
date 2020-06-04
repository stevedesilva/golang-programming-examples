package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// go run -race main.go
	funcWithAtomic()
}

func funcWithAtomic() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)

}
