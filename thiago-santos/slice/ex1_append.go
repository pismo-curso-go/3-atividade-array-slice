
// ex1_append.go
package main

import "fmt"

func main() {
  names := []string{"Ana", "Bruno", "Carla"}
  names = append(names, "Daniel", "Eva")
  fmt.Println("Names:", names)
}
