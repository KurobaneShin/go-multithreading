package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	// "sync"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		time.Sleep(500 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("you had visited this page %d times.", number)))
	})

	http.ListenAndServe(":8080", nil)
}
