package main

import (
	"fmt"
	"math"
)

func main() {
	a := 20
	b := 30

	//Suma, resta, division, multiplicacion
	fmt.Println("Sumando ", a, " + ", b, " = ", a+b)
	fmt.Println("Restando ", a, " - ", b, " = ", a-b)
	fmt.Println("Dividiendo ", a, " / ", b, " = ", a/b)
	fmt.Println("Multiplicando ", a, " * ", b, " = ", a*b)
	fmt.Println("Modulo ", a, " % ", b, " = ", a%b)

	//asignacion
	// a+=
	// a-=
	// a/=
	// a*=
	// a%=

	fmt.Println("Pi", math.Pi)
	fmt.Println("Euler", math.E)
	fmt.Println("Potencia: ", math.Pow(float64(a), float64(b)))
	fmt.Println("Raiz cuadrada: ", math.Sqrt(8))
	fmt.Println("Raiz cubica: ", math.Cbrt(27))
}
