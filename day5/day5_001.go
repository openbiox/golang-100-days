// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HandleFunc func(pattern string, handler func(ResponseWriter, *Request))
	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.
	http.HandleFunc("/", handler) // each request calls handler

	// log.Fatal is equivalent to Print() followed by a call to os.Exit(1).
	// http.ListenAndServe func(addr string, handler Handler) error
	// ListenAndServe listens on the TCP network address addr and then calls Serve with
	// handler to handle requests on incoming connections.
	// Accepted connections are configured to enable TCP keep-alives.
	// The handler is typically nil, in which case the DefaultServeMux is used.
	// ListenAndServe always returns a non-nil error.
	log.Fatal(http.ListenAndServe("localhost:8989", nil))
}

// handler echoes the Path component of the request URL r.
// like io.Writer, but for web stream
// required for handler http.ResponseWriter and http.Request
// Like Node.js Express.js req, res
// router.get('/*', function (req, res, next) {
// })
func handler(w http.ResponseWriter, r *http.Request) {
	// r contains various useful value and function
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
