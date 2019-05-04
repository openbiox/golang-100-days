package main

// line break auto convert => ;
// fmt pkg: string manuplation (input, output, paste)
// os pkg: interact with operating system (cross platform)
import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("[Debug]", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("vim-go")

	// practice fmt.Println: split by space
	fmt.Println("vim-go", "123")

	// os.Args object like Python [1,2,3], R c(1,2,3)
	// os.Args[0]: get the program name
	fmt.Println(os.Args[0])

	// auto set the s and sep to ''
	var s, sep string
	// := (short variable declaration) just for undefiend variable
	// and auto declare according to the given value (e.g. int, float, string, bool)
	// # Print all command line parameters
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
		// string plus (similar with Python plus) like R paste function
		// Equal to s = s + sep + os.Args[i]
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
