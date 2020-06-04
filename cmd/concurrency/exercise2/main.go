package main

import "fmt"


type person struct {
	first string
}


type human interface {
	speak() string
}

func (p *person) speak() string {
	return fmt.Sprintf("hello %s", p.first)
}



func saySomething(h human) {
	fmt.Println("Human is speaking ", h.speak())
}
// Summary
// Inteface method is implemented if:
// receiver and value match 
// receiver and pointer value
// Not implmented if:
// ptr receiver and value
func main() {
	p := person{"steve"}

	// fmt.Println("Can call speak directly on type... ",p.speak())

	// NOTE : This rule only applies to interfaces
	// you CANNOT pass a value of type person into saySomething
	// saySomething(p)

	// you CAN pass a value of type *person into saySomething
	saySomething(&p)

}
