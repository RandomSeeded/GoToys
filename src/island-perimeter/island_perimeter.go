package main

import "fmt"

func calculatePerimeter(grid [][]int) int64 {
  var visited map[string]int
  visited = make(map[string]int)
  visited["0,0"] = 1
  
  fmt.Println("visited",visited)
  return 1
}

func secondFunction() {
  fmt.Printf("This is a second function")
}

func main() {
  fmt.Printf("Hello world!")
  secondFunction()

  grid := make([][]int, 3)
  for i := 0; i < 3; i++ {
    grid[i] = make([]int, 3)
  }
  grid[0][0] = 1
  grid[1][0] = 1
  grid[1][1] = 1
  fmt.Println("grid", grid)

  fmt.Println("output", calculatePerimeter(grid))
}
