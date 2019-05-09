## flag

For commandline flag parse.

```go
var n = flag.Bool("n", false, "omit trailing newline")

// flag.String func(name string, value string, usage string) *string
// flag.String defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag.
// must use *sep to use the real value
var sep = flag.String("s", " ", "separator")

func main() {
	// flag.Parse is required before to use the real flag values
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
```

## new

`new` can be overwrited. 

```go
// newInt equal newInt2 func
func newInt() *int {
	return new(int)
}
func newInt2() *int {
	var dummy int
	return &dummy
}
```

## type and interface

Struct is clear. `func (type) String() string` can be used for `fmt.Println` print string.

```go
// defein temperature types ;
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	// fmt.Println(c == f)       // ompile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// for fmt.Println
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```