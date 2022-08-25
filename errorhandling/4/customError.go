package main

import "fmt"

type man struct {
	name string
	age  int
	err  error
}

func main() {
	_, err := sqrt(-3)
	if err != nil {
		fmt.Println(err)
	}

}

func (n man) Error() string {
	return fmt.Sprintln(n.name, n.age, n.err)
}

func sqrt(v int) (int, error) {
	if v < 0 {
		var e error = fmt.Errorf("The value is below acceptable level")
		return 0, man{"Aj", 22, e}
	}
	return 3, nil
}
