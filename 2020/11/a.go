package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func move(seats []string) ([]string, bool) {
	newSeats := make([]string, len(seats))
	changed := false
	for y := 0; y < len(seats); y++ {
		for x := 0; x < len(seats[y]); x++ {
			adjCnt := 0
			for rx := -1; rx <= 1; rx++ {
				for ry := -1; ry <= 1; ry++ {
					if rx == 0 && ry == 0 {
						continue
					}
					nx, ny := x+rx, y+ry
					if ny < 0 || ny >= len(seats) || nx < 0 || nx >= len(seats[y]) {
						continue
					}
					if byte(seats[ny][nx]) == '#' {
						adjCnt++
					}
				}
			}
			if byte(seats[y][x]) == 'L' && adjCnt == 0 {
				newSeats[y] += "#"
				changed = true
			} else if byte(seats[y][x]) == '#' && adjCnt >= 4 {
				newSeats[y] += "L"
				changed = true
			} else {
				newSeats[y] += string(seats[y][x])
			}
		}
	}
	return newSeats, changed
}

func main() {
	seats := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		seats = append(seats, scanner.Text())
	}
	for {
		newSeats, changed := move(seats)
		if !changed {
			break
		}
		seats = newSeats
	}
	for _, s := range seats {
		log.Println(s)
	}

	result := 0
	for y := 0; y < len(seats); y++ {
		for x := 0; x < len(seats[y]); x++ {
			if seats[y][x] == '#' {
				result++
			}
		}
	}
	fmt.Println(result)
}
