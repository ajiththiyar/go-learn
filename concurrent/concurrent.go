package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("GOOS:\t\t", runtime.GOOS)
	fmt.Println("GOArch:\t\t", runtime.GOARCH)
	fmt.Println("GoCPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
	wg.Add(1)
	go bar()
	foo()

	fmt.Println("GoCPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
}

func bar() {
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
