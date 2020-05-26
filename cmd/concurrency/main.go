package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*Hands-on exercise #1
in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
code: https://github.com/GoesToEleven/go-programming
video: 148
*/
var wg sync.WaitGroup

const value int = 5
func main() {

	fmt.Printf("Starting CPU %d \n", runtime.NumCPU())
	fmt.Printf("Starting Go routines %d \n",runtime.NumGoroutine())
	MainLoop()

	wg.Add(2)
	go LoopOne()
	go LoopTwo()
	fmt.Printf("MID CPU %d \n", runtime.NumCPU())
	fmt.Printf("MID Go routines %d \n",runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("End CPU %d \n", runtime.NumCPU())
	fmt.Printf("End Go routines %d \n",runtime.NumGoroutine())
}

func MainLoop() {
	for i := 0; i < value; i++ {
		fmt.Printf("Loop Main \tn# %d \n",i)
	}
	fmt.Println()
}

func LoopOne() {
	for i := 0; i < value; i++ {
		fmt.Printf("Loop One \tn# %d \n",i)
	}
	fmt.Println()
	wg.Done()
}

func LoopTwo() {
	for i := 0; i < value; i++ {
		fmt.Printf("Loop Two  \tn# %d \n",i)
	}
	fmt.Println()
	wg.Done()
}