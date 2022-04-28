package main

import "fmt"

func main() {
  var (
    x int
    p int
    y int
    n int = 0
  )
  fmt.Scan(&x, &p, &y)
  for {
    x += x * p / 100
    if x >= y {
      break
    }
    n++
  }
  fmt.Println(n + 1)

}