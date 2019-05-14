package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	const e = 2.71828
	fmt.Println(e)

	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	fmt.Println(Avogadro, Planck) // 6.02214129e+23 6.62606957e-34

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

	fmt.Println(math.Min(56, 22)) // return 22
}
