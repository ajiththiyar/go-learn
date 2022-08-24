package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	bs, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	fmt.Println("...............................................")

	error := bcrypt.CompareHashAndPassword(bs, []byte("password123"))
	if error != nil {
		fmt.Println("Login failed")
	} else {
		fmt.Println("Login Success")
	}
}
