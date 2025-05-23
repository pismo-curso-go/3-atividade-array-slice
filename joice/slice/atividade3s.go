package main

import "fmt"

func main() {
	numeros := [5]int{10, 20, 30, 40, 50}
	soma := 0

	for _, num := range numeros {
		soma += num
	}

	fmt.Println("Total:", soma)
}
