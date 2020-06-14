package main

import "fmt"

func main() {
	func1()
	func2()
}

func func1() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

}

func func2() {
	cr := make(chan int)

	go func(ch chan<- int) {
		ch <- 42
	}(cr)
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
