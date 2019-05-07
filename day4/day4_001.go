// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// http from net/http
		resp, err := http.Get(url)
		if err != nil {
			// Fprintf func(w io.Writer, format string, a ...interface{}) (n int, err error)
			// Fprintf formats according to a format specifier and writes to w.
			// It returns the number of bytes written and any write error encountered.
			// output stand error message (shell pipe 2> error.msg)
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)

			// exit this program
			os.Exit(1)
		}
		// ReadAll func(r io.Reader) ([]byte, error)
		// ReadAll reads from r until an error or EOF and returns the data it read.
		// A successful call returns err == nil, not err == EOF.
		// Because ReadAll is defined to read from src until EOF,
		// it does not treat an EOF from Read as an error to be reported.
		// resp.Body is a readable server response flow (as the io.Reader type)
		b, err := ioutil.ReadAll(resp.Body)

		// close io.Reader
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
