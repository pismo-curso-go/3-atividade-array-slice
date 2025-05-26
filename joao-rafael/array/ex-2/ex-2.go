package main

import "fmt"

func main() {
	numbers := [4]int{1, 2, 3, 4}
	sum := 0

	var input int
	fmt.Println("insert a number: ")
	fmt.Scan(&input)

	for _, number := range numbers {
		sum += number
	}

	fmt.Println("sum: ", sum)
}
