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
	x, y, d := 0, 0, 0
	for _, ist := range ists {
		switch com, num := getInstruction(ist); com {
		case 'E':
			x += dx[0] * num
			y += dy[0] * num
		case 'N':
			x += dx[1] * num
			y += dy[1] * num
		case 'W':
			x += dx[2] * num
			y += dy[2] * num
		case 'S':
			x += dx[3] * num
			y += dy[3] * num
		case 'L':
			d += num / 90
			d %= 4
		case 'R':
			d -= (num / 90) % 4
			d = (d + 4) % 4
		case 'F':
			x += dx[d] * num
			y += dy[d] * num
		}
		log.Print(x, " ", y)
	}
	fmt.Println(abs(x) + abs(y))
}
