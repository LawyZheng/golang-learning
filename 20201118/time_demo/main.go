package main

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

func printNow() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}

func countTime() {
	start := time.Now()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
	}
	end := time.Now()
	fmt.Println((end.Sub(start)).Microseconds())

}

func getCaller() {
	pc, file, line, ok := runtime.Caller(1)
	fmt.Println(pc, file, line, ok)
	fmt.Println(path.Base(file))
}

func main() {
	// printNow()
	getCaller()

}
