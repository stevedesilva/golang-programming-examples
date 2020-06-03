package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	funcWithRaceCondition()
	// funcWithMutex()
	// funcWithAtomic()

}

// example of a race condition
// go run -race main.go
func funcWithRaceCondition() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// same as
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)

}

func funcWithMutex() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	var mu sync.Mutex
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter

			// time.Sleep(time.Second * 1)
			// same as
			runtime.Gosched()

			v++
			counter = v

			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)

}

func funcWithAtomic() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	var (
		counter int64
		wg      sync.WaitGroup
	)

	const gs = 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			time.Sleep(time.Second * 1)
			// runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)

}
