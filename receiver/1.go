package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Printf("%s is of %d years \n", p.name, p.age)
}

func (p person) printPointer() {
	fmt.Printf("Ikkinchi: %s is of %d years \n", p.name, p.age)
}

func main() {
	alex := person{
		name: "Alex",
		age:  18,
	}
	alex.print()
	alex.printPointer()
}
