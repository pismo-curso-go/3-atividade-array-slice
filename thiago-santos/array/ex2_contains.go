
// ex2_contains.go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  arr := [4]int{10, 20, 30, 40}
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter a number: ")
  text, _ := reader.ReadString('\n')
  n, _ := strconv.Atoi(strings.TrimSpace(text))

  found := false
  for _, v := range arr {
    if v == n {
      found = true
      break
    }
  }
  if found {
    fmt.Printf("%d is in %v\n", n, arr)
  } else {
    fmt.Printf("%d is NOT in %v\n", n, arr)
  }
}
