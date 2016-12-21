package main

import "fmt"

func calculatePerimeter(grid []int) int64 {
  fmt.Printf("Aaaaah")
  return 1
}

func secondFunction() {
  fmt.Printf("This is a second function")
}

func main() {
  fmt.Printf("Hello world!")
  secondFunction()

  grid := [5]int{1,2,3,4,5}
  calculatePerimeter(grid)
}
