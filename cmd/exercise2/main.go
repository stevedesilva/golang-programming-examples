package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() string {
	return fmt.Sprintf("hello %s", p.first)
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println("Human is speaking ", h.speak())
}

func main() {
	p := person{"steve"}

	// fmt.Println("Can call speak directly on type... ",p.speak())

	// you CANNOT pass a value of type person into saySomething
	// saySomething(p)

	// you CAN pass a value of type *person into saySomething
	saySomething(&p)

}
