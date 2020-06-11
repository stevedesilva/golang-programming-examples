package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutInWithLimit(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")

}

func populate(c chan int) {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	defer close(c2)
	wg := sync.WaitGroup{}
	const maxgoroutines = 10

	for v := range c1 {
		fmt.Println("fanOutIn>>", v)
		wg.Add(1)

		// This code would result in timeconsumingWork receiving the last value of v (e.g. 9),
		// so we must pass in the current value so the func will take a copy of the value
		// go func(){
		// 	c2 <- timeconsumingWork(v)
		// 	wg.Done()
		// }()

		go func(v int) {
			c2 <- timeconsumingWork(v)
			wg.Done()
		}(v)

	}

	wg.Wait()
}

func fanOutInWithLimit(c1 <-chan int, c2 chan<- int) {
	defer close(c2)
	wg := sync.WaitGroup{}

	const maxgoroutines = 10
	wg.Add(maxgoroutines)

	for i := 0; i < maxgoroutines; i++ {
		// launch 10 go routines
		go func() {
			// listen on c1 which blocks until a message is received
			for v := range c1 {
				// run the work and put unto a channel
				c2 <- timeconsumingWork(v)
			}
			// exit when c1 closed
			wg.Done()
		}()
	}

	wg.Wait()
}

func timeconsumingWork(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	fmt.Println("timeconsumingWork>>", v)
	return v + rand.Intn(1000)
}
