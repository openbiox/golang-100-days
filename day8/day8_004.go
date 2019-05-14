package main

import (
	"fmt"
	"unicode/utf8"
)

// btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob reports whether i is non-zero.
func itob(i int) bool { return i != 0 }

func main() {
	s := "hello, world"
	s1 := []string{"a", "b"}
	fmt.Println(len(s)) // "12"
	// fmt.Println(cap(s)) not work, only for []string{"a", "b"}
	fmt.Println(cap(s1))    // 2
	fmt.Println(len(s1))    // 2
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')

	fmt.Println(s[0:5]) // "hello"

	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	unis := "Hello, 世界"
	fmt.Println(unis)      // Hello, 世界
	fmt.Println(len(unis)) // 13

	/*
		"世界"
		"\xe4\xb8\x96\xe7\x95\x8c"
		"\u4e16\u754c"
		"\U00004e16\U0000754c"
	*/
	// RuneCount func(p []byte) int
	// RuneCount returns the number of runes in p. Erroneous and short encodings are treated as single runes of width 1 byte.
	// utf8.RuneCount

	// RuneCountInString is like RuneCount but its input is a string.
	fmt.Println(utf8.RuneCountInString(unis)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for i := 0; i < len(unis); {
		r, size := utf8.DecodeRuneInString(unis[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	/*
		0	h
		1	e
		2	l
		3	l
		4	o
		5	,
		6
		7	w
		8	o
		9	r
		10	l
		11	d

		0	H
		1	e
		2	l
		3	l
		4	o
		5	,
		6
		7	世
		10	界
	*/

	// range for string
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
