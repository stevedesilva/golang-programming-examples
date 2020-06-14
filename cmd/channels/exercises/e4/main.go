package main

import (
	"fmt"
)

// pull the values off the channel using a select statement
func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 100; i++ {
			c <- i
		}
		fmt.Println("gen: Quit")
		q <- 0
		fmt.Println("gen: Close")
	}()

	return c
}

func receive(c, quit <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-quit:
			fmt.Println("Quit receive")
			return
		}
	}
}

