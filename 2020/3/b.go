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

	stepX := []int{1, 3, 5, 7, 1}
	stepY := []int{1, 1, 1, 1, 2}
	result := 1
	for i, _ := range stepX {
		cnt := 0
		x := 0
		for y := 0; y < len(Map); y += stepY[i] {
			if Map[y][x] == '#' {
				cnt++
			}
			x += stepX[i]
			x %= len(Map[y])
		}
		result *= cnt
	}

	fmt.Println(result)
}
