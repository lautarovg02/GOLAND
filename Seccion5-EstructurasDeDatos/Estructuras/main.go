package main

import "fmt"

/*
* Una struct es una estructura de datos
* Una Struct seria el modelo de una clase
* que permite definir un tipo de dato personalizado
* compuesto por diferentes campos con nombres y tipos especificicos
 */
type Persona struct {
	nombre string
	edad   int
	correo string
}

/*
* Creando un metodo. para que esta func sea un metodo necesitamso hacignarle un recepetor.
* Este se coloca antes del nombre, en este caso entre parentesis, GO nos recomienda
* Que pasemos un puntero, osea la referencia de memoria.
* Con el receptor indicamos de que Struct es el metodo.
 */

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es: ", p.nombre)
}

func main() {
	//creando datos de tipo persona

	/*
		*Como asignar valores?
		*OPCION 1:
			var feli Persona
			feli.nombre = "Felipe"
			feli.edad = 28
			feli.correo = "feli@gmai.com"
		*OPCION 2:
			feli := Persona{"Alex", 28, "feli@gmail.com"}
	*/

	feli := Persona{"Feli", 28, "feli@gmail.com"}
	lau := Persona{"Lau", 38, "lau@gmail.com"}

	/*
		 *Metodos
		* La unica diferencia con las funciones, es que estas son independientes,
		* En cambio, los metodos pertenecen a una struct asignando un receptor
	*/
	fmt.Println(feli, lau)
	fmt.Println(feli.nombre)
	lau.diHola()
	/*
	* Punteros
	* Uso : Los punteros son utilizados como receptores de metodos en GO, permite modificar directamente la variable original
	* Puntero : Es una variable que almacena la direccion de la memoria de otra variable
	 */
	//Con & + el bombre de la variable podemos acceder a la referencia de memoria
	var x int = 10
	var y int = 10
	/*Funcion para modificar valor de una variable*/
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)

	/*Funcion que intenta modificar el valor de una variable, pero no puede
	* Por que en ves de pasar la misma direccion de memoria, estamos pasando una copia
	 */
	fmt.Println(y)
	editarSinPuntero(y)
	fmt.Println(y)

}

// funcion que recibe un puntero como parametro
func editar(x *int) {
	*x = 20
}
func editarSinPuntero(x int) {
	x = 20
}
