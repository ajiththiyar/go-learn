package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type animal struct {
	breed    string
	strength int
}

type being interface {
	speak()
}

func foo(b being) {
	b.speak()
}

func main() {
	p := person{
		"aj",
		23,
	}

	a := animal{
		"chicken",
		59,
	}
	foo(p)
	foo(a)
	foo(person{"ar", 18})

	fmt.Println("..........................................................")
	SwitchOnType(34)
	SwitchOnType("dsakljf")
	fmt.Println("..........................................................")
	x := boo()
	fmt.Printf("%v\n", x())
	fmt.Printf("%T\n", boo())

}

func boo() func() int {
	return func() int {
		return 69
	}
}

func (p person) speak() {
	fmt.Println("I am ", p.name, "and age is", p.age)
}

func (a animal) speak() {
	fmt.Println("I am ", a.breed, "and strength is", a.strength)
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}
