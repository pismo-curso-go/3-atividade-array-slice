package main

import "fmt"

func main() {
	originalNameList := []string{
		"John",
		"Doe",
		"Mary",
	}
	originalNameList = addNameToList("Jane", originalNameList)
	originalNameList = addNameToList("Moe", originalNameList)

	fmt.Println("Name list:")
	for _, name := range originalNameList {
		fmt.Printf("> %s\n", name)
	}

	fmt.Println("------------------------------")

	originalNumberList := []int{1, 2, 3, 4, 5}
	fmt.Println("Removendo posição 2 do slice:")
	fmt.Println(originalNumberList)

	newNumberList := removePositionFromList(2, originalNumberList)
	fmt.Println("Resultado:")
	fmt.Println(newNumberList)
}

func addNameToList(newName string, list []string) []string {
	return append(list, newName)
}

func removePositionFromList(idx int, list []int) []int {
	if idx < 0 || idx >= len(list) {
		fmt.Println("Indice fora dos limites, não foi possivel remover")
		return list
	}
	return append(list[:idx], list[idx+1:]...)
}
