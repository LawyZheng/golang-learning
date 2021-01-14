package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	num    int64
	result int64
	worker int
}

func newJob(num int64) *Job {
	return &Job{num: num, result: int64(0), worker: -1}
}

func createNum(jobChan chan<- *Job) {
	r := rand.NewSource(time.Now().UnixNano())
	// for i := 0; i < jobs; i++ {
	// 	jobChan <- r.Int63()
	// }
	for {
		job := newJob(r.Int63())
		jobChan <- job
	}
	close(jobChan)
}

func calNum(i int, resultChan chan<- *Job, jobChan <-chan *Job) {
	for job := range jobChan {
		job.worker = i
		var tmp int64
		num := job.num
		for num > 0 {
			tmp += (num % 10)
			num = num / 10
		}
		job.result = tmp
		resultChan <- job
	}
}

var once sync.Once

func main() {
	jobChan := make(chan *Job)
	resultChan := make(chan *Job, 100)

	go createNum(jobChan)

	for i := 0; i < 24; i++ {
		go calNum(i, resultChan, jobChan)
	}

	// for i := 0; i < jobs; i++ {
	// <-resultChan
	// tmp := <-resultChan
	// fmt.Println("result = ", tmp)
	// }

	// for {
	// 	<-resultChan
	// }

	for job := range resultChan {
		fmt.Printf("Num = %d, result = %d. Calculated By Workder %d.\n", job.num, job.result, job.worker)
	}
}
