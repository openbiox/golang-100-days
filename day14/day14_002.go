package main

// sort key
import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"charlie": 34,
		"alice":   31,
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	/*
		[charlie alice]
		[alice charlie]
		alice	31
		charlie	34
	*/
}
