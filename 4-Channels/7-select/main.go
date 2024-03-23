package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id      int64
	message string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var id int64 = 0

	go func() {
		for {
			time.Sleep(time.Second * 3)
			c1 <- Message{id, "Message for channel 1"}
			atomic.AddInt64(&id, 1)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- Message{id, "Message for channel 2"}
			atomic.AddInt64(&id, 1)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")
		}
	}
}
