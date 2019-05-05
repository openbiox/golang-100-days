// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	// bufio can be used to read keyboard input
	"bufio"
	"fmt"
	"os"
)

func main() {
	// store key-value pairs ("map" object)
	// e.g. python dict: {"a": 1, "b": 2}
	// e.g. R list: obj$a = 1; obj$a = 2
	// "make" create Golang "map" object
	counts := make(map[string]int)

	// bufio.NewScanner: read data stream with cache IO, and require io.Reader (e.g. stringtrings.NewReader("a b c"), and os.Stdin) as parameter
	// os.Stdin: standard input
	input := bufio.NewScanner(os.Stdin)

	// iterate all lines
	// input.Scan can auto remove \n
	for input.Scan() {
		// input.Text() get the line content
		counts[input.Text()]++
		// equal to
		// line := input.Text()
		// counts[line] = counts[line] + 1
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		// n > 1: require line is recurrent
		if n > 1 {
			// %d number, %s string, \n newline
			// print various lines counts (unique lines)
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
