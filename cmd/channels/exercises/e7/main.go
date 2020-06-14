package main

import "fmt"

/*
write a program that:
	launches 10 goroutines
	each goroutine adds 10 numbers to a channel
	pull the numbers off the channel and print them

*/
const grMax = 10

func main() {
	out := make(chan int)
	done := make(chan int)
	in := addData(done)

	go printNumbers(in, out, done)

	for d := range out {
		fmt.Println(d)
	}
	fmt.Println("Exiting!")
}

func printNumbers(in, out, done chan int) {
	var doneCount int
	defer close(out)
	defer close(done)
	defer close(in)
	for {
		select {
		case d := <-in:
			fmt.Println("Received ", d)
			out <- d
		case v := <-done:
			doneCount += v
			if doneCount == grMax {
				fmt.Println("Closing channels!",doneCount)
				return

			}

		}
	}

}

func addData(done chan int) chan int {
	data := make(chan int)
	go func() {
		// launches 10 goroutines
		for i := 1; i <= grMax; i++ {
			go func(j int) {
				for i := 0; i < 10; i++ {
					data <- i + j
				}
				done <- 1
			}(i)
		}
	}()

	return data

}
