package main

import (
	"fmt"
)

// Show the comma ok idiom
func main() {
	bufferedOk()
	iifeOk()
}

func bufferedOk() {
	c := make(chan int, 1)

	c <- 1
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func iifeOk() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
