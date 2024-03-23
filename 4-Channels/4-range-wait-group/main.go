package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	go publish(channel)
	go reader(channel, &waitGroup)
	waitGroup.Wait()
}

func reader(channel chan int, waitGroup *sync.WaitGroup) {
	for i := range channel {
		fmt.Printf("Received %d\n", i)
		waitGroup.Done()
	}
}

func publish(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
}
