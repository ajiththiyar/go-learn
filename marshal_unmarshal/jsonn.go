package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
}

func main() {

	p1 := person{"aj", "atk"}
	p2 := person{"ar", "atk"}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {

		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

	people1 := []person{}

	er := json.Unmarshal(bs, &people1)

	if err != nil {
		fmt.Println(er)
	}
	fmt.Printf("%+v", people1)

}
