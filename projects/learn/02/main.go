package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Calculadora de linea de comandos")
	fmt.Println("Formato: número operador número (por ejemplo 5 + 3)")

	fmt.Print("Ingrese la operacion: ")
	expresion, _ := reader.ReadString('\n')

	expresion = strings.TrimSpace(expresion)

	parts := strings.Split(expresion, " ")
	if len(parts) != 3 {
		fmt.Println("Formato incorrecto. Debe ser número operador número.")
		return
	}

	numero1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Error al convertir el primero numero", err)
		return
	}

	numero2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Error al convertir el segundo numero:", err)
		return
	}

	resultado := 0.0

	switch parts[1] {
	case "+":
		resultado = numero1 + numero2

	case "-":
		resultado = numero1 - numero2

	case "*":
		resultado = numero1 * numero2

	case "/":
		if numero2 != 0 {
			resultado = numero1 / numero2
		} else {
			fmt.Println("Operador no valido. Use +, -, *")
			return
		}
	default:
		fmt.Println("Operador no valido")
		return
	}
	fmt.Println("RESULADO: %.2f\n", resultado)
}
