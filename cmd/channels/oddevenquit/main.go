package main

import "fmt"

func main() {

	odd := make(chan int)
	eve := make(chan int)
	quit := make(chan int) 

	go send(eve, odd,quit)

	receive(eve,odd,quit)

	fmt.Printf("exit from main\n")
}

func receive(eve,odd,quit <-chan int) {
	for {
		select {
		case v:= <- eve : 
			fmt.Printf("from eve channel:\t %v\n",v)
		
		case v:= <- odd : 
			fmt.Printf("from odd channel:\t %v\n",v)
		case v,ok:= <- quit : 
			if !ok {
				fmt.Printf("from quit channel, closing:\t %v\n",v)
				return
			} 
			fmt.Printf("from quit channel:\t %v\n",v)
		
		} 
	}
}

func send(eve,odd,quit chan<- int) {
	for i:=0;i <100;i++ {
		if i % 2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
	close(quit)
}