package main

import (
	"fmt"
	"time"
)

func main() {
	firstToReturnChannel()
	fmt.Printf("-----------------------\n")
	waitForAll()
}

func waitForAll() {
	c := make(chan int)
	numThreads := 4

	for i := 0; i < numThreads; i++ {
		go runRandomTasks(i, c)
	}
	for i := 0; i < numThreads; i++ {
		fmt.Printf("Thread number %v has completed\n", <-c)
	}
}

func firstToReturnChannel() {
	c := make(chan int)
	numThreads := 4

	for i := 0; i < numThreads; i++ {
		go runRandomTasks(i, c)
	}
	completedThread := <-c
	// expect to only receive 1
	fmt.Printf("First thread to complete: %v\n", completedThread)
}

func runRandomTasks(threadNumber int, c chan int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Thread number: %d\n", threadNumber)
	c <- threadNumber
}
