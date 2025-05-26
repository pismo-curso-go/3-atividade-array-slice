package main

import "fmt"

func main() {

	//task 1
	fmt.Println("Array Example")
	intArray := [5]int{10, 7, 28, -5, 8}
	fmt.Println("Array:", intArray)
	fmt.Println("Sum of array elements:", arraySum(intArray))

	//task 2
	fmt.Println("------------------------------------------")
	intArray2 := [4]int{1, 2, 3, 4}
	fmt.Printf("Array: %d\n", intArray2)
	fmt.Println("Insert a number to check if it is present in the array:")
	var value int
	fmt.Scanf("%d", &value)
	if checkIsPresent(intArray2, value) {
		fmt.Printf("The number %d is present in the array.\n", value)
	} else {
		fmt.Printf("The number %d is not present in the array.\n", value)
	}
}

func arraySum(intArray [5]int) int {
	sum := 0
	for _, value := range intArray {
		sum += value
	}
	return sum
}

func checkIsPresent(intArray [4]int, value int) bool {
	for _, v := range intArray {
		if v == value {
			return true
		}
	}
	return false
}
