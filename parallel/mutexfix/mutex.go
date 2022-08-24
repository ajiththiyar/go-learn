package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("GoCPU", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	const gs = 100
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(gs)
	count := 0
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := count
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			count = v
			m.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println("GoRoutines", runtime.NumGoroutine())
}
