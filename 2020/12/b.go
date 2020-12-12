package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func toInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
		return -1
	} else {
		return i
	}
}

func getInstruction(s string) (byte, int) {
	return byte(s[0]), toInt(s[1:])
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func main() {
	ists := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ists = append(ists, scanner.Text())
	}

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	wx, wy, wd := 10, 1, 0
	x, y := 0, 0
	for _, ist := range ists {
		switch com, num := getInstruction(ist); com {
		case 'E':
			wx += dx[0] * num
			wy += dy[0] * num
		case 'N':
			wx += dx[1] * num
			wy += dy[1] * num
		case 'W':
			wx += dx[2] * num
			wy += dy[2] * num
		case 'S':
			wx += dx[3] * num
			wy += dy[3] * num
		case 'L':
			wd += num / 90
			wd %= 4
		case 'R':
			wd -= (num / 90) % 4
			wd = (wd + 4) % 4
		case 'F':
			x += wx * num
			y += wy * num
		}
		switch wd {
		case 1:
			wx, wy = -wy, wx
		case 2:
			wx, wy = -wx, -wy
		case 3:
			wx, wy = wy, -wx
		}
		wd = 0
		log.Print(x, " ", y)
	}
	fmt.Println(abs(x) + abs(y))
}
