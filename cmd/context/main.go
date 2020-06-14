package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for v := range data(ctx) {
		if v == 5 {
			return
		}
		fmt.Println(v)
	}
	fmt.Println("Exiting")
}

func data(ctx context.Context) <-chan int {
	out := make(chan int)
	// n := 1
	go func() {
		for {
			select {
			// case out <- n:
			// 	fmt.Println("rec n: ", n)
			// 	n++

			case <-ctx.Done():
				fmt.Println("Done")
			case t := <-time.After(time.Microsecond * 1000 ):
				fmt.Println("time", t)
				// default:
				// 	fmt.Println("default")
			}
		}
	}()
	return out
}
