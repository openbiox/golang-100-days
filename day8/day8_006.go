package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	// bytes、strings、strconv, unicode, strings
	fmt.Println(basename("/tmp/kk.txt")) // kk

	fmt.Println(basename2("/tmp/cc.a")) // cc

	fmt.Println(path.Base("/tmp/kk.txt")) // kk.txt
	fmt.Println(path.Dir("/tmp/kk.txt"))  // /tmp

	fmt.Println(comma("accc"))    // a,ccc
	fmt.Println(comma("acccccc")) // a,ccc,ccc

	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s, b, s2) // abc [97 98 99] abc
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
