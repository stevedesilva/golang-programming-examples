package main

import (
	"fmt"
)

func main() {
	buffer()
	iife()
}

func buffer() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

// Immediately-invoked Function Expression
func iife() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
