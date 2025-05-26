package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Sum of Elements
	var array [5]int
	fmt.Println("-------------------------------------------------")
	fmt.Println("Exercise 1: Sum of Elements")
	fmt.Println("-------------------------------------------------")
	fmt.Println("Enter 5 integers:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	sum := 0
	for _, v := range array {
		sum += v
	}
	fmt.Println("Sum of elements:", sum)

	// Exercise 2: Check if a number is present
	fmt.Println("-------------------------------------------------")
	fmt.Println("Exercise 2: Check if a number is present")
	fmt.Println("-------------------------------------------------")
	var num int

	fmt.Print("Enter a number to search in the array: ")
	fmt.Scan(&num)
	found := false
	for _, v := range array {
		if v == num {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Number found in the array!")
	} else {
		fmt.Println("Number not found in the array.")
	}
}
