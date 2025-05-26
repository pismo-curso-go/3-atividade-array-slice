package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"joao", "rafael", "mateus"}

	names = append(names, "bianca", "vera")

	fmt.Println("names: ", strings.Join(names, ", "))
}
