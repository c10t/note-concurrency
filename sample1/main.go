package main

import (
  "fmt"
  "sync"
)

func main() {
  fmt.Println("Hello, World!")

  var wg sync.WaitGroup
  salutation := "hello"

  wg.Add(1)

  go func() {
    defer wg.Done()
    salutation = "welcome"
  }()

  wg.Wait()
  fmt.Println(salutation)
}
