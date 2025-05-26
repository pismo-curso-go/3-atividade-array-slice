package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	numbers = append(numbers[:2], numbers[2+1:]...)

	fmt.Println("new slice: ", numbers)
}
