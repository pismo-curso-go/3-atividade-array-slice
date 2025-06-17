package main

import (
	"fmt"
)

func main() {
	
	nomes := []string{"Pedro", "Boris", "Luna"}

	
	nomes = append(nomes, "Nina", "Eduardo")



	
	numeros := []int{1, 2, 3, 4, 5}

	
	pos := 2
	numeros = append(numeros[:pos], numeros[pos+1:]...)

	
	fmt.Println("Slice após remover o elemento da posição 2:", numeros)

}
