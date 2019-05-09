// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

// flag.Bool func(name string, value bool, usage string) *bool
// flag.Bool defines a bool flag with specified name, default value, and usage string. The return value is the address of a bool variable that stores the value of the flag.
// must use *n to use the real value
var n = flag.Bool("n", false, "omit trailing newline")

// flag.String func(name string, value string, usage string) *string
// flag.String defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag.
// must use *sep to use the real value
var sep = flag.String("s", " ", "separator")

func main() {
	// flag.Parse is required before to use the real flag values
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
