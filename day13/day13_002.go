package main

import "fmt"

// example of golang slice
func main() {
	months := [...]string{1: "January", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	// not allow > cap(s)
	// fmt.Println(summer[:20])    // panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"

	intArray := []int{1: 1, 2: 2, 3: 3, 4: 4}
	fmt.Println(intArray) // [0 1 2 3 4]
	reverse(intArray)
	fmt.Println(intArray) // [4 3 2 1 0]
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
