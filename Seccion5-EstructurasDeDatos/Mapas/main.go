package main

import "fmt"

func main() {
	//como definir un mapa
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	// iterar un map
	for clave, valor := range colors {
		fmt.Printf("\n clave: %s, valor:%s", clave, valor)
	}

	fmt.Println(colors)
	//acceder a un elemetno mediante la clase
	fmt.Println(colors["rojo"])
	//agregar un nuevo elemento
	colors["negro"] = "#000000"
	fmt.Println(colors)
	//como recuperar un valor
	valorRojo := colors["rojo"]
	fmt.Println("Valor color rojo: ", valorRojo)
	// aparte de recueprar el dato podemos imprimir una verificacion

	if valor, ok := colors["verde"]; ok {
		fmt.Println("Existe la clave", valor)
	} else {
		fmt.Println("No existe la clave")
	}
	// como eliminar un elemento
	delete(colors, "negro")
	fmt.Println(colors)
}
