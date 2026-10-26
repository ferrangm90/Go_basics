package main

import "fmt"

// ejercicio para practicar el uso de variables, condicionales, switch, bucles y entrada y salida de datos por teclado

func main() {
	var examenes, comportamiento, trabajos float64
	var notaExamenes, notaComportamiento, notaTrabajos float64
	var texto string

	// Porcentajes de la nota final
	for {
		fmt.Println("\nPORCENTAJES (deben sumar 100) o pulsa 's' para salir")
		fmt.Println("----------------------------------------------------\n")

		// Pedir examenes
		fmt.Print("Porcentaje de examenes: ")
		fmt.Scan(&texto)
		if texto == "s" || texto == "S" {
			fmt.Println("Hasta pronto")
			return
		}
		fmt.Sscan(texto, &examenes)

		// Pedir comportamiento
		fmt.Print("Porcentaje de comportamiento: ")
		fmt.Scan(&texto)
		if texto == "s" || texto == "S" {
			fmt.Println("Hasta pronto")
			return
		}
		fmt.Sscan(texto, &comportamiento)

		// Pedir trabajos
		fmt.Print("Porcentaje de trabajos: ")
		fmt.Scan(&texto)
		if texto == "s" || texto == "S" {
			fmt.Println("Hasta pronto")
			return
		}
		fmt.Sscan(texto, &trabajos)

		// Comprobar que estan entre 0 y 100
		if examenes < 0 || examenes > 100 || comportamiento < 0 || comportamiento > 100 || trabajos < 0 || trabajos > 100 {
			fmt.Println("Los porcentajes deben ser números entre 0 y 100\n")
			continue
		}

		// Comprobar que suman 100
		suma := examenes + comportamiento + trabajos
		if suma == 100 {
			break
		}

		fmt.Printf("Los porcentajes suman %.0f%% en lugar de 100%%. Vuelve a introducirlos.\n\n", suma)
	}

	// Puntuaciones
	fmt.Println("\nPUNTUACIONES (0 a 10) o pulsa 's' para salir")
	fmt.Println("--------------------------------------------\n")


	// Nota examenes
	for {
		fmt.Print("Nota de exámenes (0-10): ")
		fmt.Scan(&texto)
		if texto == "s" || texto == "S" {
			fmt.Println("Hasta pronto")
			return
		}
		fmt.Sscan(texto, &notaExamenes)
		if notaExamenes >= 0 && notaExamenes <= 10 {
			break
		}
		fmt.Println("La nota debe estar entre 0 y 10.")
	}

	// Nota comportamiento
	for {
		fmt.Print("Nota de comportamiento (0-10): ")
		fmt.Scan(&texto)
		if texto == "s" || texto == "S" {
			fmt.Println("Hasta pronto")
			return
		}
		fmt.Sscan(texto, &notaComportamiento)
		if notaComportamiento >= 0 && notaComportamiento <= 10 {
			break
		}
		fmt.Println("La nota debe estar entre 0 y 10.")
	}

	// Nota trabajos
	for {
		fmt.Print("Nota de trabajos (0-10): ")
		fmt.Scan(&texto)
		if texto == "s" || texto == "S" {
			fmt.Println("Hasta pronto")
			return
		}
		fmt.Sscan(texto, &notaTrabajos)
		if notaTrabajos >= 0 && notaTrabajos <= 10 {
			break
		}
		fmt.Println("La nota debe estar entre 0 y 10.")
	}

	// Bonus o penalización por entrega de trabajos
	fmt.Println("\nEntrega de Trabajos")
	fmt.Println("-------------------\n")
	fmt.Println("1: Entrega excelente (+1 punto)")
	fmt.Println("2: Entrega normal (sin cambios)")
	fmt.Println("3: Entrega con retraso (-1 punto)")
	fmt.Println("4: Entrega incompleta (-2 puntos)")

	for {
		fmt.Print("Elige una opción (1-4) o 's' para salir: ")
		fmt.Scan(&texto)

		if texto == "s" || texto == "S" {
			fmt.Println("Programa finalizado.")
			return
		}

		opcionValida := true

		switch texto {
		case "1":
			notaTrabajos = notaTrabajos + 1
			fmt.Println("Bonus aplicado: +1 punto.")
		case "2":
			fmt.Println("Sin cambios en la nota de trabajos.")
		case "3":
			notaTrabajos = notaTrabajos - 1
			fmt.Println("Penalización aplicada: -1 punto.")
		case "4":
			notaTrabajos = notaTrabajos - 2
			fmt.Println("Penalización aplicada: -2 puntos.")
		default:
			fmt.Println("Opción no válida. Debe ser 1, 2, 3 o 4.")
			opcionValida = false
		}

		if opcionValida {
			break
		}
	}

	// Controlamos que la los puntos que suman y restan segun la entrega no se salgan del rango 0-10
	if notaTrabajos > 10 {
		notaTrabajos = 10
	}
	if notaTrabajos < 0 {
		notaTrabajos = 0
	}
	fmt.Printf("Nota final de trabajos: %.2f\n\n", notaTrabajos)

	// Nota final
	notaFinal := (notaExamenes*examenes + notaComportamiento*comportamiento + notaTrabajos*trabajos) / 100

	fmt.Println("\nResumen")
	fmt.Println("---------")
	fmt.Printf("Exámenes (%.0f%%): %.2f\n", examenes, notaExamenes)
	fmt.Printf("Trabajos (%.0f%%): %.2f\n", trabajos, notaTrabajos)
	fmt.Printf("Comportamiento (%.0f%%): %.2f\n", comportamiento, notaComportamiento)
	fmt.Println("---------\n")
	fmt.Printf("Nota final: %.2f\n", notaFinal)

	// Aprovado o suspendido
	if notaFinal >= 5 {
		fmt.Println("APROBADO")
	} else {
		fmt.Println("SUSPENDIDO")
	}
}