package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, error := os.Open("name.txt")
	if error != nil {
		fmt.Println(error)
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
