package main

import "fmt"

func main() {
	ages := make(map[string]int) // mapping from strings to ints
	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	// equal to
	ages = make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	delete(ages, "alice") // remove element ages["alice"]

	// ages["bob"]  return 0
	ages["bob"] = ages["bob"] + 1 // happy birthday!

	fmt.Println(ages)

	//_ = &ages["bob"] // compile error: cannot take address of map element

	// names := make(map[string]string)
	// names := map[string]string not work
	names := map[string]string{
		"lastname":  "Jianfneg",
		"firstname": "Li",
	}
	fmt.Println(names)
}
