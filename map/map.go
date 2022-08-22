package main

import "fmt"

func main() {
	m := map[string]int{"misspenny": 22, "bond": 34}

	fmt.Println(m)

	fmt.Println(m["bond"])
	// return 0 when not present
	fmt.Println(m["none"])

	print(m, "bond")

	print(m, "misspenny")
	m["null"] = 36
	print(m, "null")

	fmt.Println("Iterating through loop")

	loop(m)

	delete(m, "bond")
	fmt.Println("After deleting bond")
	loop(m)
}

func loop(m map[string]int) {
	for i, v := range m {
		fmt.Println(i, v)
	}
}

func print(m map[string]int, b string) {
	// returns value as well as present or not ; ok to check if it is present wont print if ok is false i.e. not present in map
	if v, ok := m[b]; ok {
		fmt.Println(b, v, ok)
	}
}
