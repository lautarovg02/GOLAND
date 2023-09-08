package main

import "fmt"

func main() {
	name := "Lautaro"
	age := 20

	//Funciones de Fmt
	//print imprime
	fmt.Print("Hola")
	//println imprime con un salto de linea
	fmt.Println("Holaaaaa")
	//Printf = podemos parsear informacion en una cadena
	fmt.Printf("Hola me llamo %s y tengo %d .\n", name, age)
	greeting := fmt.Sprintf("Hola me llamo %s y tengo %d .\n", name, age)
	fmt.Println(greeting)

	//Funciones de entrada de Datos
	var nombre, apellido string
	var edad int
	fmt.Println("Ingrese su nombre: ")
	// Scanln podemos guardar datos ingresados desde el teclado, recibe la referencia de memoria de una variable
	fmt.Scanln(&nombre, &apellido)
	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&edad)
	//%v usamos cuando no sabemos el tipo de dato que vamos a colocar
	fmt.Printf("Su nombre es %s %s y tiene %d \n", nombre, apellido, edad)
	//%T para saber que tipo de dato contiene la variable
	fmt.Printf("El tipo de dato nombre es: %T \n", nombre)
	fmt.Printf("El tipo de dato edad es: %T", edad)

}
