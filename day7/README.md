## action scope

```go
func f() {}
var g = "g"

func main() {
	f := "f"
	fmt.Println(f) // "f"; local var f shadows package-level func f
	fmt.Println(g) // "g"; package-level var
	fmt.Println(h) // compile error: undefined: h

	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	fmt.Println(x, y) // compile error: x and y are not visible here

	if f, err := os.Open(fname); err != nil { // compile error: unused: f
		return err
	}
	f.ReadByte() // compile error: undefined f
	f.Close()    // compile error: undefined f

	// Right way
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	f.ReadByte()
	f.Close()

}
```

## init()

Initiation function for Golang package vars.

```go
	var Cwd string

	// If import from external, the Cwd will be updated as the current workdir
	func init() {
		var err error
		cwd, err = os.Getwd()
		if err != nil {
			log.Fatalf("os.Getwd failed: %v", err)
		}
	}
```

## Calulate

```go
* / % << >> & &^
+ - | ^
== != < <= > >=
&&
||
```

## Reverse Loop

```go
medals := []string{"gold", "silver", "bronze"}
for i := len(medals) - 1; i >= 0; i-- {
	fmt.Println(medals[i]) // "bronze", "silver", "gold"
}
```