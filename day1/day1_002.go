package main

// line break auto convert => ;
// fmt pkg: string manuplation (input, output, paste)
// os pkg: interact with operating system (cross platform)
import (
	"fmt"
	"os"
	"strings"
)

func PlusLoop() {

	// initial the var s and sep using short variable declaration (only exists in function)
	// s := ""
	// equal to :
	// var s string
	// var s = ""
	// var s string = ""
	s, sep := "", ""

	// use another way get all parameters
	// range returen key-value pairs
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		// string plus (similar with Python plus) like R paste function
		// Equal to s = s + sep + os.Args[i]
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
func JoinDemo() {
	// strings.Join do same thing (recommand for big size input)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
