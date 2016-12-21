package main

import "fmt"

func calculatePerimeter(grid [][]int) int64 {
  visited := make(map[string]int)

  var startingX int
  var startingY int
  for i := 0; i < len(grid); i++ {
    var breaking bool
    for j := 0; j < len(grid[i]); j++ {
      fmt.Println(i, j)
      if grid[i][j] == 1 {
        startingX = i
        startingY = j
        fmt.Println("starting points set")
        breaking = true
        break
      }
    }
    if breaking == true {
      break
    }
  }

  fmt.Println("startingX", startingX)
  fmt.Println("startingY", startingY)
  
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
