package main

import "fmt"

// thread 1
func main() {
	channel := make(chan string)

	// thread 2
	go func() {
		channel <- "Hello" // it blocks until thread 1 reads from the channel
	}()

	// thread 1
	msg := <-channel
	fmt.Println(msg)
}
