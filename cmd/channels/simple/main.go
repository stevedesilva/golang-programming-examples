package main

import "fmt"



func main() {
	ch := make(chan int)
	go sender(ch)
	
	// main routine blocks - ela
	reader(ch)
}

func sender (ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch <-chan int) {
	// reads and waits until close()
	for i := range ch {
		fmt.Println(i)
	}
	
}
