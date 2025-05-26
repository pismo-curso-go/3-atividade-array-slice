package main

import "fmt"

func main() {
	numbers := [4]int{1, 2, 3, 4}

	var input int
	fmt.Print("insert a number: ")
	fmt.Scan(&input)

	found := false
	for _, number := range numbers {
		if number == input {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("the number was not found in the array")
		return
	}

	fmt.Println("the number was found in the array")
}
