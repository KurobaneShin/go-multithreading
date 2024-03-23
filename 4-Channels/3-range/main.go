package main

import "fmt"

func main() {
	channel := make(chan int)
	go publish(channel)
	reader(channel)
}

func reader(channel chan int) {
	for i := range channel {
		fmt.Printf("Received %d\n", i)
	}
}

func publish(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}
