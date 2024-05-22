package main

import (
	"fmt"
)

func main() {
	var ganancias, perdidas, taxes float64 = 5.99, 3.0, 1.1

	fmt.Print("Ganancias: ")
	fmt.Scan(&ganancias)
	fmt.Print("Perdidas: ")
	fmt.Scan(&perdidas)

	// calculo de ganancias
	gananciasNetas := ganancias - perdidas
	// calculo de impuestos
	impuestos := gananciasNetas * taxes
	// calculo de ganancias netas

	var gananciasTotales = gananciasNetas - impuestos

	fmt.Println("Ganancias Totales: ", gananciasTotales)
}
