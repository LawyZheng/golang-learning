package main

import (
	"fmt"
)

// var wg sync.WaitGroup
// var ch1 chan int
// var ch2 chan int

func a(ch1, ch2 chan int) {
	// defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x * x
	}
	close(ch2)
}

func b(ch1 chan int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go a(ch1, ch2)
	go a(ch1, ch2)
	go b(ch1)

	for i := range ch2 {
		fmt.Println(i)
	}

	// wg.Wait()
}
