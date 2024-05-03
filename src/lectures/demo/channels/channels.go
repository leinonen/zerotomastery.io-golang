package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

// <-chan = read only channel
// chan<- = write only channel
func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {

	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("unhandled control message")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {

	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {

		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timed out")
			control <- DoExit
			<-control
			fmt.Println("Program exit")
			return
		}
	}
}
