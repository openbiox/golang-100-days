package main

import "fmt"

// Package tempconv performs Celsius and Fahrenheit temperature computations.

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
