package main

import "fmt"

func main() {
	// Atividade 1
	nomes := []string{"Alice", "Bob", "Carlos"}
	nomes = append(nomes, "Diana", "Eduardo")
	fmt.Println("Nomes:", nomes)

	// Atividade 2
	numeros := []int{1, 2, 3, 4, 5}
	posicaoRemover := 2
	
	numeros = append(numeros[:posicaoRemover], numeros[posicaoRemover+1:]...)
	fmt.Println("Números após remoção:", numeros)
}
