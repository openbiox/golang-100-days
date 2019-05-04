package main

// line break auto convert => ;
// fmt pkg: string manuplation (input, output, paste)
// os pkg: interact with operating system (cross platform)
import (
	"strings"
	"testing"
)

func BenchmarkPlusLoop(b *testing.B) {
	s1 := make([]string, b.N)
	s2 := ""
	for n := 0; n < b.N; n++ {
		s2 = s2 + s1[n]
	}
}

func BenchmarkJoinLoop(b *testing.B) {
	s := make([]string, b.N)
	strings.Join(s, "\n")
}
