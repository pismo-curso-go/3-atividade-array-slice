package remove

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)

	fmt.Println("Slice after removing element at index 2:", numbers)
}
