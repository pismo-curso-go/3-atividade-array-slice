package main

import "fmt"

func main() {
	numeros := [4]int{10, 20, 30, 40}

	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	encontrado := false
	for _, n := range numeros {
		if n == num {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("Número encontrado!")
	} else {
		fmt.Println("Número NÃO encontrado.")
	}
}
