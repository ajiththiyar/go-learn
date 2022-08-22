package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretagent struct {
	p  person
	id string
}

func main() {
	m := person{
		first: "james", last: "bond", age: 34,
	}
	fmt.Println(m.age, m.first, m.last)
	print(m)

	sa := secretagent{
		p:  m,
		id: "hancock",
	}

	fmt.Println(sa.id, sa.p.age, sa.p.first, sa.p.last)

	fmt.Println(sa)

	sa.speak()
}

func (s secretagent) speak() {
	fmt.Println("Hello my name is", s.p.first, s.p.last)
}

func print(p person) {
	fmt.Println(p)
}
