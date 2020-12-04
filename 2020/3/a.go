package main

import (
  "fmt"
  "log"
)

func main() {
  var Map []string
  for {
    var line string
    _, err := fmt.Scan(&line)
    if err != nil {
      log.Print(err)
      break
    }

    Map = append(Map, line)
  }

  x := 0
  result := 0
  for _, L := range Map {
    x %= len(L)
    if L[x] == '#' {
      result++
    }
    x += 3
  }
  
  fmt.Println(result)
}

