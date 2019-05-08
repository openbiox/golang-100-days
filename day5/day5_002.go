// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// type Mutex struct {
//   state    int32
//   sema    uint32
// }
// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
// A Mutex must not be copied after first use.
//
// To avoid different goroutine concorrent modified the count value
var mu sync.Mutex
var count int

func main() {
	// if access /count see 0, then access / and then /count is 3
	// access /count will activiate two goroutine for handler and counter
	// so see /count is 0, the real value of count is 1 now
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
