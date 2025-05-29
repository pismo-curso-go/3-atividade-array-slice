package main

import "fmt"

func main() {
	names := []string{"Daniel", "Alice", "David"}
	fmt.Println("First names:", names)

	names = append(names, "Débora", "Carlos")
	fmt.Println("Names after append:", names)
}
