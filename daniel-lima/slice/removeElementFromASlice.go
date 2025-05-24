package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 2; i < len(numbers)-1; i++ {
		numbers[i] = numbers[i+1]
	}

	// Reduce the slice
	numbers = numbers[:len(numbers)-1]

	fmt.Println("Slice after removal (position 2):", numbers)
}
