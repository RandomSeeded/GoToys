package main

import "fmt"

func calculatePerimeter(grid [][]int) int64 {
  return 1
}

func secondFunction() {
  fmt.Printf("This is a second function")
}

func main() {
  fmt.Printf("Hello world!")
  secondFunction()

  grid := make([][]int, 5)
  for i := 0; i < 5; i++ {
    grid[i] = make([]int, 5)
    for j := 0; j < 5; j++ {
      grid[i][j] = i + j
    }
  }
  fmt.Println("grid", grid)

  fmt.Println("output", calculatePerimeter(grid))
}
