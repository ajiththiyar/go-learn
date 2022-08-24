package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("GoCPU", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var count uint64
	for i := 0; i < gs; i++ {
		go func() {
			//time.Sleep(time.Second)
			runtime.Gosched()
			atomic.AddUint64(&count, 1)
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println("GoRoutines", runtime.NumGoroutine())
}
