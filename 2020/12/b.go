package main

import (
	"../helper"
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInstruction(s string) (byte, int) {
	return byte(s[0]), helper.ToInt(s[1:])
}

func main() {
	var ists []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ists = append(ists, scanner.Text())
	}

	di := map[byte]int{'E': 0, 'N': 1, 'W': 2, 'S': 3}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	wx, wy, wd := 10, 1, 0
	x, y := 0, 0
	for _, ist := range ists {
		switch com, num := getInstruction(ist); com {
		case 'E', 'N', 'W', 'S':
			wx += dx[di[com]] * num
			wy += dy[di[com]] * num
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
	fmt.Println(helper.Abs(x) + helper.Abs(y))
}
