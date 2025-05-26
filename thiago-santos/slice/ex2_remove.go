
// ex2_remove.go
package main

import "fmt"

func main() {
  nums := []int{1, 2, 3, 4, 5}
  // remover o elemento na posição 2 (valor 3)
  nums = append(nums[:2], nums[3:]...)
  fmt.Println("After removal:", nums)
}
