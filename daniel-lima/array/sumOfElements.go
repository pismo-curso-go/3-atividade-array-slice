package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("The sum of the elements is: %d\n", sum)
}
