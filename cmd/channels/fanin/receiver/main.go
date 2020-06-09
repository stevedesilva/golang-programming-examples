package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(eve, odd)

	go receive(eve, odd, fanin)

	for v := range fanin {
		fmt.Printf("FANIN\t %d\n", v)
	}

	fmt.Println("About to exit.")
}

func send(eve, odd chan<- int) {
	defer close(eve)
	defer close(odd)

	for i := 0; i < 10; i++ {
		if even := i%2 == 0; even {
			eve <- i
		} else {
			odd <- i
		}
	}
}

// general to specfic chan
func receive(eve, odd <-chan int, fanin chan<- int) {
	defer close(fanin)

	var wg sync.WaitGroup // struct init to {{} [0 0 0]}
	

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range eve {
			fmt.Printf("EVEN\t %d\n", i)
			fanin <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := range odd {
			fmt.Printf("ODD\t %d\n", i)
			fanin <- i
		}
	}()

	// waiting until go routines are finished then close fanin channel
	wg.Wait()
}
