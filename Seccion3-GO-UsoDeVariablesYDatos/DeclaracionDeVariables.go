package main

import (
	"fmt"
)

var nameCourse = "golanddd"

func main() {
	//* VAR : Se puede declarar tanto fuera de un funcion como dentro
	//* Para importar varias paquetes usamos el import mas parentesis, como abajo.
	//Seccion3-Declaracion De Variables
	fmt.Println("Nombre del curso: " + nameCourse)

	var name, schol string
	var numberPhone int

	//Asignar Valores
	name = "Lautaro"
	schol = "Secundaria 10"
	numberPhone = 2023134

	//utilizar variables
	fmt.Println(name, schol, numberPhone)

	//Declaracion de variables de difernte tipo
	//var (
	//	pepe, pepe2 string
	//	number                   int
	//)
	//fmt.Println(pepe, pepe2, number)
	/*
	* CUANDO INICIALIZAMOS UNA VARIABLE DE UNA,
	* NO ES NECESARIO PONER EL TIPO
	* EJEMPLO var name string = "carlos" ES IGUAL QUE var name = "carlos"
	 */

	//var (
	//	pepe   = "pepe"
	//	pepe2  = "pepe2"
	//	number = 3
	//)

	/*
	*EL TIPO DE DECLARACION DE ARRIBA ES LA MISMA QUE LA DE ABAJO
	 */
	// de acuerdo a la posicion de las variables asignamos el valor
	var pepe, pepe2, number = "pepe", "pepe2", 14

	fmt.Println(pepe, pepe2, number)

	//Declaracion de variables e incializar a la ves
	var (
		firstName  string = "Feli"
		secondName string = "Rivarola"
		age        int    = 13
	)

	//utilizar variables
	fmt.Println(firstName, secondName, age)
	fmt.Println("Nombre: "+firstName+" "+secondName+" edad: ", age)

	// := se utiliza para declarar una variable nueva e inicializarlas
	// solo podemos usar este tipo de declaracion dentro de una funcion
	city, adress := "New York", "Liberty 4323"
	fmt.Println("City: " + city + " adress: " + adress)

}
