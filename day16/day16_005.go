package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}
type St string

func (s St) Describe() {
	fmt.Println("envoked!")
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	case string:
		fmt.Println("String var")
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
	st := St("mystring")
	findType(st)
	findType(p)
}
