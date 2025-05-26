package main

import "fmt"

func main() {

	//Atividade 1
	fmt.Println("Exercício 1:")
	array := [5]int{1, 2, 3, 4, 5}
	var total int

	for _, number := range array {
		total = total + number
	}
	fmt.Printf("Total: %d", total)
	fmt.Println("------------------------------------")

	//Exercícioos 2
	fmt.Println("Exercício 2")

	array2 := [4]int{1, 2, 3, 4}
	var input int
	fmt.Println("Digite um número")

	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}

	var hasNumber bool
	for i, _ := range array2 {
		if array2[i] == input {
			hasNumber = true
		}
	}

	if hasNumber == true {
		fmt.Printf("O número %d está presente", input)
	} else {
		fmt.Printf("O número %d não está presente", input)
	}
}
