package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// Metodo para agregar Tarea
func (l *ListaTareas) agregarTarea(tarea Tarea) {
	l.tareas = append(l.tareas, tarea)
}

// Metodo para marcar tarea como completada
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

func (l *ListaTareas) editarTarea(index int, tarea Tarea) {
	l.tareas[index] = tarea
}

func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	//Instancia lista tareas
	lista := ListaTareas{}

	/* *Bucle infinito*/
	/* *Instancia de bufio para entrada de datos */
	//Creamos un nuevo lector que implementa un os.Stdin, esta variable representa la entrada standar del lector del teclado
	// con esto podemos leer muchos caracteres ya que con fmt tenemos limitaciones
	/* *ReadString devuelve dos valores, la cadena y un error*/
	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opcion:\n" +
			"1. Agregar tarea\n" +
			"2. Marcar tarea como completada\n" +
			"3. Editar tarea\n" +
			"4. Eliminar tarea\n" +
			"5. Salir")

		fmt.Print("Ingrese la opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente.")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que desea actualizar: ")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea eliminar : ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			fmt.Println("Saliendo del programa...")
			return // termina de ejecutar main
		default:
			fmt.Println("Opcion invalida.")
		}

		//Listar todas las tareas
		fmt.Println("Lista de tareas: ")
		fmt.Println("=============================================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("=============================================================")

	}

}
