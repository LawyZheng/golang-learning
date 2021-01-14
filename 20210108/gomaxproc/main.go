package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	now := time.Now().Unix()
	fmt.Println(now)
	fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(4)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
