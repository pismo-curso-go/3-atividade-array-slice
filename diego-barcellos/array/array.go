package main

import "fmt"

func main() {

	fmt.Println("Array Example")
	intArray := [5]int{10, 7, 28, -5, 8}
	fmt.Println("Array:", intArray)
	fmt.Println("Sum of array elements:", arraySum(intArray))


}

func arraySum(intArray [5]int) int {
	sum := 0
	for _, value := range intArray {
		sum += value
	}
	return sum
}