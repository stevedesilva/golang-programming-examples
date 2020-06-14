package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// funcWithRaceCondition()
	// funcWithChanSolution()
	funcWithChanAndSyncSolution()
}

// example of a race condition
// go run -race main.go
func funcWithRaceCondition() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	// shared mem
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

func funcWithChanSolution() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	ch := make(chan int)

	const gs = 100
	for i := 0; i < gs; i++ {
		go func() {
			runtime.Gosched()
			// increment by 1
			ch <- 1
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())

	}

	var counter int
	// add result
	for i := 0; i < gs; i++ {
		counter += <-ch
	}
	close(ch)

	fmt.Println("Counter: ", counter)

}

func funcWithChanAndSyncSolution() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	inCh := make(chan int)
	//outCh := make(<-chan int)
	//defer close(outCh)

	//
	const gs = 100
	for i := 0; i < gs; i++ {
		go func() {
			runtime.Gosched()
			// increment by 1
			inCh <- 1
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	defer close(inCh)

	var counter int
	// add result
loop:
	for {
		select {
		case v := <-inCh:
			counter += v
		case <-time.After(time.Millisecond * 1000):
			break loop
		}

	}

	fmt.Println("Counter: ", counter)

}
