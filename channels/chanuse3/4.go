package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(eve, odd, quit)
	receive(eve, odd, quit)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even", v)
		case v := <-o:
			fmt.Println("Odd", v)
		case v := <-q:
			fmt.Println("EXIT TRiggered", v)
			return
		}
	}

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
