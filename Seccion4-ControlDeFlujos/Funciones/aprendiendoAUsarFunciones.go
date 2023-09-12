package main

import "fmt"

func main() {
	hello()
	greeting := helloName("lautaro")
	fmt.Println(greeting)
	sum, mul := calc(4, 5)
	fmt.Println("La suma es: ", sum)
	fmt.Println("La mul es: ", mul)
}
func sumar(a, b int) int {
	return a + b
}

//	func calc(a, b int) (int, int) {
//		sum := a + b
//		mul := a * b
//		return sum, mul
//	}
//
// esto es lo mismo que el de arriba
func calc(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}

func hello() {
	fmt.Println("Hola desde la funcion ")
}

func helloName(name string) string {
	return "Hola " + name
}
