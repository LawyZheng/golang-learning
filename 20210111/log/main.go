package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func startLogging(logChan <-chan string) {
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("Open File Error: ", err)

	}
	defer file.Close()
	for log := range logChan {
		file.Write([]byte(log))
	}

}

var logChan chan string

var wg sync.WaitGroup

func main() {
	logChan = make(chan string)

	go startLogging(logChan)

	i := 0
	for {
		s := "This is log for " + strconv.Itoa(i) + "\n"
		logChan <- s
		i++
		time.Sleep(500 * time.Millisecond)
	}

	// createLog()

}
