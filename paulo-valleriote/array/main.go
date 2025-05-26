package main

import "fmt"

func arrSum(arr [5]int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func arrHasNum(target int, arr [4]int) bool {
	isPresent := false
	for _, n := range arr {
		if target == n {
			isPresent = true
			break
		}
	}
	return isPresent
}

func main() {
	elementsToSum := [5]int{8, 6, 3, 7, 1}
	sum := arrSum(elementsToSum)
	fmt.Println("A soma é igual a:", sum)

	var targetNumber int
	fmt.Println("-----------------")
	fmt.Print("Qual numero deseja verificar? > ")
	fmt.Scanf("%d", &targetNumber)
	arrToSearch := [4]int{9, 4, 3, 2}
	isPresent := arrHasNum(targetNumber, arrToSearch)
	if isPresent {
		fmt.Printf("O número {%d} está presente", targetNumber)
		return
	}
	fmt.Printf("O número {%d} não está presente", targetNumber)
}
