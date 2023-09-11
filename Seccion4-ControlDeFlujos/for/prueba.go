package main

import "fmt"

func main() {
	/*
		for {
		fmt.Println("Bucle infinito")
		}*/

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	/*
	*Break: Dentro de los bucles For en Go
	* la palabra break es clave. Se utiliza para salir
	* antes de la condicion de corte
	 */
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
	/*
	*Continue: Se utiliza para saltar a la sig iteracion
	* de un bucle, sin ejecutar el codigo despues de esta
	* instruccion "Continue"
	 */
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
