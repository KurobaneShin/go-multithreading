package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for i := range data {
		fmt.Printf("Worker %d received %d\n", workerId, i)
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan int)

	worketQuantity := 1000000

	for i := 0; i < worketQuantity; i++ {
		go worker(i, channel)
	}

	for i := 0; i < 10000000; i++ {
		channel <- i
	}
}
