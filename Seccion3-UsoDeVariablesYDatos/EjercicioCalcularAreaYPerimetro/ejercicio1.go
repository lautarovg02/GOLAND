package main

import (
	"fmt"
	"math"
)

/*
*Ejercicio: Calcule e imprima el área y el perímetro del triángulo
*Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule
*e imprima el área y el perímetro del triángulo.
*El programa debe:
*Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.

	*Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.
	*Calcular el área del triángulo usando la fórmula base x altura / 2.
	*Calcular el perímetro del triángulo sumando los lados.
	*Imprimir el área y el perímetro del triángulo con dos decimales de precisión.
	*El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro.
	*También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.
	*Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.
*/
const MAXDECIMALES = 2

func main() {

	var longitudLadoA, longitudLadoB float64

	//entrada de datos
	fmt.Println("ingrese lado 1: ")
	fmt.Scanln(&longitudLadoA)
	fmt.Println("ingrese lado 2: ")
	fmt.Scanln(&longitudLadoB)
	fmt.Println("Lado 1: ", longitudLadoA, " Lado 2: ", longitudLadoB)

	//Calculos
	area := (longitudLadoA * longitudLadoB) / 2
	hipotenusa := math.Sqrt(math.Pow(longitudLadoA, 2) + math.Pow(longitudLadoB, 2))
	perimetro := longitudLadoA + longitudLadoB + hipotenusa
	fmt.Printf("\nArea: %.*f", MAXDECIMALES, area)
	fmt.Printf("\nPerimetro: %.*f", MAXDECIMALES, perimetro)
}
