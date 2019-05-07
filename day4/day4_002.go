// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}
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

		_, err = io.Copy(os.Stdout, resp.Body)

		// close io.Reader
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
