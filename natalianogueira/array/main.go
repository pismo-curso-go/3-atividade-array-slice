package main

import "fmt"

func main() {
   
    numeros1 := [5]int{2, 19, 14, 6, 4}
    soma := 0

    for _, valor := range numeros1 {
        soma += valor
    }

    fmt.Println("ATIVIDADE 1 PROPOSTA  1")
    fmt.Println("A soma dos elementos do array é:", soma)

    numeros2 := [4]int{5, 12, 7, 20}
    var procurado int

    fmt.Println("\n ATIVIDADE 2 PROPOSTA 2")
    fmt.Print("Digite um número: ")
    fmt.Scan(&procurado)

    encontrado := false
    for _, valor := range numeros2 {
        if valor == procurado {
            encontrado = true
            break
        }
    }

    if encontrado {
        fmt.Println("Número esta presente no array.")
    } else {
        fmt.Println("O Número nao esta no array.")
    }
}
