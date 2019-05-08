// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		switch k {
		case "num":
			{
				val, _ := strconv.Atoi(v[0])
				if val%2 == 0 {
					fmt.Fprintf(w, "%d is not odd\n", val)
				} else {
					fmt.Fprintf(w, "%d is odd\n", val)
				}
			}
		default:
			fmt.Println("landed on edge!")
		}
	}
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
