package main

import (
	"fmt"
)

/*
write a program that
puts 100 numbers to a channel
pull the numbers off the channel and print them

*/

type doneCh chan struct{}
type dataCh chan int

func main() {
	done := make(doneCh)

	dataCh := createData(done)

	receiveData(dataCh, done)

	fmt.Println("Exit program")

}

func receiveData(dataCh dataCh, done doneCh) {
	for {
		select {
		case v := <-dataCh:
			fmt.Println(v)
		case <-done:
			fmt.Println("That's it, I'm done!")
			return
		}
	}
}

func createData(done doneCh) dataCh {
	dataCh := make(dataCh)
	go func() {
		defer close(dataCh)
		// some work
		for i := 0; i < 100; i++ {
			dataCh <- i
		}

		// done <- struct{}{}
		close(done)

	}()
	return dataCh

}
