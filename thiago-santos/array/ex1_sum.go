
// ex1_sum.go
package main

import "fmt"

func main() {
  arr := [5]int{1, 2, 3, 4, 5}
  sum := 0
  for _, v := range arr {
    sum += v
  }
  fmt.Printf("Sum of %v = %d\n", arr, sum)
}
