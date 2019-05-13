package main

import (
	"fmt"

	"github.com/openbiox/golang-100-days/day7/biogolang"
)

func main() {
	var i biogolang.People

	i.Name = "Jianfeng"
	i.Age = 20
	i.Gender = "Male"
	fmt.Println(i)
	fmt.Println(biogolang.Cwd)

	v := biogolang.Num8 + 127
	fmt.Println(biogolang.Unmu8, v, v+1, v*v) // 0 127 -128 1

	var k uint8
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Println(k, x, y) // 0 34 6
	// x<<n for x * 2^n
	fmt.Println(x<<1, x>>1) // 68 17

	fmt.Printf("%08b\n", k) // "00000000"
	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	// reverse loop

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}

	f := 3.141 // a float64
	m := int(f)
	fmt.Println(f, m) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"

	fnm := 1e100 // a float64
	m = int(fnm)
	fmt.Println(m)

	// %[1] re-use %d; %# for %o / %x / %X print 0 / 0x or 0X profix
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	z := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", z)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"

}
