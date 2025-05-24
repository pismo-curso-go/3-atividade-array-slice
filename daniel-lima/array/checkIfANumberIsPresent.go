package main

import "fmt"

func main() {
	numbers := [4]int{15, 25, 35, 45}
	var search int

	fmt.Print("Enter a number to check: ")
	fmt.Scan(&search)

	found := false
	for _, num := range numbers {
		if num == search {
			found = true
			break
		}
	}

	if found {
		fmt.Printf("The number %d is present in array.\n", search)
	} else {
		fmt.Printf("The number %d is NOT present in the array.\n", search)
	}
}
