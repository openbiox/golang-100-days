## math.Exp math.NaN math.Inf math.Min

```go
    for x := 0; x < 8; x++ {
		// Printf for format control
		// math.Exp is R 2^x
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
		/*
			x = 0 e^x =    1.000
			x = 1 e^x =    2.718
			x = 2 e^x =    7.389
			x = 3 e^x =   20.086
			x = 4 e^x =   54.598
			x = 5 e^x =  148.413
			x = 6 e^x =  403.429
			x = 7 e^x = 1096.633
		*/
    }
    
    nan := math.NaN()
    fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
    
    inf := math.Inf
    fmt.Println(inf == nil) // "false"
    // inf > nil not work

    fmt.Println(math.Min(56, 22)) // return 22
```

## SVG header

```go
// SVG header
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
```

## utf8

```go
unis := "Hello, 世界"
fmt.Println(unis)      // Hello, 世界
fmt.Println(len(unis)) // 13

fmt.Println(utf8.RuneCountInString(unis)) // "9"

s := "プログラム"
	// % x
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	fmt.Println(string(r)) // プログラム

	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"

	fmt.Println(string(1234567)) //�
```

## path

```go
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
```

## strings

```go
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
```

## bytes

```go
// string => slice
func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) [][]byte
func HasPrefix(s, prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte

// buffer is useful
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	// WriteByte func(c byte) error
	// WriteByte appends the byte c to the buffer, growing the buffer as needed. The returned error is always nil, but is included to match bufio.Writer's WriteByte. If the buffer becomes too large, WriteByte will panic with ErrTooLarge.
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
```