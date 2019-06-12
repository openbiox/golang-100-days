## Questions

You're given strings `J` representing the types of stones that are jewels, and `S` representing the stones you have.  Each character in `S` is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in `J` are guaranteed distinct, and all characters in `J` and `S` are letters. Letters are case sensitive, so `"a"`is considered a different type of stone from `"A"`.

Example 1:

Input: J = "aA", S = "aAAbbbb"
Output: 3

Example 2:

Input: J = "z", S = "ZZ"
Output: 0

Note:

-   `S` and `J` will consist of letters and have length at most 50.
-   The characters in `J` are distinct.

## My solutions

```golang
package main

import "strings"

func main(){
    J := "aAAAaaa111"
    S := "aAAA"
    fmt.Printlin(numJewelsInStones1(J, S))
}

func numJewelsInStones1(J string, S string) int {
	count := 0
	for _, i := range S {
		for _, j := range J {
			if i == j {
				count++
			}
		}
	}
	return count
}

func numJewelsInStones2(J string, S string) int {
	Jarray := []rune(J)
	Sarray := []rune(S)
	count := 0
	for i := 0; i < len(Sarray); i++ {
		for j := 0; j < len(Jarray); j++ {
			if Sarray[i] == Jarray[j] {
				count++
			}
		}
	}
	return count
}

func numJewelsInStones3(J string, S string) int {
	count := 0
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(J); j++ {
			if S[i] == J[j] {
				count++
			}
		}
	}
	return count
}

func numJewelsInStones4(J string, S string) int {
	count := 0
	keyMap := make(map[rune]uint8)
	for _, v := range J {
		keyMap[v]++
	}
	for _, v := range S {
		for k, s := range keyMap {
			if v == k {
				count += int(s)
			}
		}
	}
	return count
}

func numJewelsInStones5(J string, S string) int {
	var count uint8
	keyMap := make(map[rune]uint8)
	for _, v := range J {
		keyMap[v]++
	}
	for _, v := range S {
		for k, s := range keyMap {
			if v == k {
				count += s
			}
		}
	}
	return int(count)
}

func numJewelsInStones6(J string, S string) int {
	count := 0
	for _, v := range S {
		if strings.Contains(J, string(v)) {
			count += 1
		}
	}
	return count
}
```
