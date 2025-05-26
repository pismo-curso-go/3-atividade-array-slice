package main

import "fmt"

func main() {
	// Exercise 1: List of Names with Append
	names := []string{"Leandra", "Litman", "Sirlene"}
	fmt.Println("Initial names:", names)
	names = append(names, "Paola", "Leandro")
	fmt.Println("Updated names:", names)

	// Exercise 2: Remove Element from Slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial numbers:", numbers)
	position := 2
	numbers = append(numbers[:position], numbers[position+1:]...)
	fmt.Println("Updated numbers:", numbers)
}
