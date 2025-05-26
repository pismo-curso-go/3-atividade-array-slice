package main

import "fmt"

func main() {

	names := []string{"Diego", "Joel", "Dennis"}

	//Task 3
	fmt.Println(names)
	fmt.Println("Please, insert two more names to the slice:")
	var name1, name2 string
	fmt.Print("Name 1: ")
	fmt.Scanf("%s", &name1)
	fmt.Scanln()
	fmt.Print("Name 2: ")
	fmt.Scanf("%s", &name2)

	names = append(names, name1, name2)
	fmt.Println("Updated slice:", names)

	//Task 4
	fmt.Println("------------------------------------------")
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %d", intSlice)

	intSlice = append(intSlice[:1], intSlice[2:]...)
	fmt.Println("\nSlice after removing the second element:", intSlice)
}
