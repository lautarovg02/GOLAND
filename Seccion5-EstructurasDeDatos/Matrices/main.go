package main

import "fmt"

func main() {
	// matriz declarada sin inicializar
	var matriz [5]int
	// matriz declarada e inicializada
	var matrizInicializada = [5]int{1, 2, 3, 4, 5}
	//matriz iniciliacida pero no sabemos cuantos datos va a almacenar asique utilzzamos ...
	var matrizX = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Matriz declarada : ", matriz)
	fmt.Println("Matriz declarada e inicializada: ", matrizInicializada)
	fmt.Println("Matriz inicializada pero no sabemos cuantos datos va a almacenar: ", matrizX)
	// for comun
	for i := 0; i < len(matrizX); i++ {
		fmt.Printf("\n Elemento: %d de la posicion: %d", matrizX[i], i)
	}
	// for range
	for index, value := range matrizX {
		fmt.Printf("Index %d, valor %d\n", index, value)
	}
	// for range pero solo imprime el value
	for _, value := range matrizX {
		fmt.Printf("valor %d\n", value)
	}

	// Matriz bidimensional
	matrizBidi := [3][4]int{{1, 2, 3, 54}, {4, 5, 6, 65}, {21, 7, 8, 9}}
	fmt.Println(matrizBidi)

	// Acceder al tamaño de las filas y las columnas
	numFilas := len(matrizBidi)
	numColumnas := len(matrizBidi[0]) // Suponemos que todas las filas tienen la misma longitud

	for i := 0; i < numFilas; i++ {
		for j := 0; j < numColumnas; j++ {
			fmt.Println("fila: ", i, " col: ", j, " valor: ", matrizBidi[i][j])
		}
	}
}
