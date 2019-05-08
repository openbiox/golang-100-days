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

// handler echoes the HTTP request.
//
// GET / HTTP/1.1
// Header["Upgrade-Insecure-Requests"] = ["1"]
// Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"]
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"]
// Header["Accept-Encoding"] = ["gzip, deflate, br"]
// Header["Accept-Language"] = ["zh-CN,zh;q=0.9,en;q=0.8"]
// Header["Connection"] = ["keep-alive"]
// Host = "localhost:8000"
// RemoteAddr = "127.0.0.1:50483"
func handler(w http.ResponseWriter, r *http.Request) {
	// r.Method is GET; r.URL is /; r.Proto is HTTP/1.1
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		// For basic string printing use %s.
		// To double-quote strings as in Go source, use %q.
		// %s for q, %q for "q"
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// Parse URL query key-value
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// r.Form is like http://localhost:8000/?q=123&b=133
	// echo as below:
	// Form["q"] = ["123"]
	// Form["b"] = ["133"]
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
