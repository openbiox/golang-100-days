## Interface

From <https://studygolang.com/articles/12560?fr=sidebar>

```golang
package main

import (  
    "fmt"
)

// define interface 
type VowelsFinder interface {  
    FindVowels() []rune
}

type MyString string

// implement interface
func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

func main() {  
    name := MyString("Sam Anderson") // type conversion
    var v VowelsFinder // define interface var
    v = name 
    fmt.Printf("Vowels are %c", v.FindVowels())
    fmt.Printf("Vowels are %c", name.FindVowels())
}
```

```golang
package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}
type Contract struct {
	empId  int
	basicpay int
}
type Permanent struct {
	empId  int
	basicpay int
	jj int //Bonus
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.jj
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("total $%d", expense)
}

func main() {
	pemp1 := Permanent{1,3000,10000}
	pemp2 := Permanent{2, 3000, 20000}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
```

```golang
// empty interface for any types vars
package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {

	s := "Hello World"
	i := 55
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(s)
	describe(i)
	describe(strt)
}
```

```golang
// i.(T)
package main

import(
"fmt"
)

func assert(i interface{}){
    s:= i.(int)
    fmt.Println(s)
}

func main(){
  var s interface{} = 55
	// panic: pass string to s
  assert(s)
}
```

```golang
// i.(type)
package main

import (  
    "fmt"
)

func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("String: %s\n", i.(string))
    case int:
        fmt.Printf("Int: %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}
func main() {  
    findType("Naveen")
    findType(77)
    findType(89.98)
}
```

```golang
// compare interface and struct
package main

import "fmt"

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}

func main() {  
    findType("Naveen")
    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType(p)
}
```