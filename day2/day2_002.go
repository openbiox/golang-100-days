// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// initial a map object
	counts := make(map[string]int)
	// Get file list
	files := os.Args[1:]
	if len(files) == 0 {
		// function countLines (f *os.File, counts map[string]int) is defined at the file tail
		// Get counts of standard input lines (e.g. {"abc": 3, "a": 1})
		countLines(os.Stdin, counts)
	} else {
		// _: eat index of range
		// arg: file path
		for _, arg := range files {
			// os.Open: Open a file
			f, err := os.Open(arg)
			// print error message
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// close connected file
			f.Close()
		}
	}
	// os.Args is similar to ["a", "b", "c"] like {1: "a", "b": 2, 3: "c"}
	// "range counts" is similar to {"a": 2, "b":1, "c":3}
	// so "a" is line, 2 is n
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	// now we know bufio.NewScanner can input:
	// os.Open(filename); os.Stdin; stringtrings.NewReader("a b c")
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
