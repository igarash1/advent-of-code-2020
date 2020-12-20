package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"../helper"
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
	x, y, d := 0, 0, 0
	for _, ist := range ists {
		switch com, num := getInstruction(ist); com {
		case 'E', 'N', 'W', 'S':
			x += dx[di[com]] * num
			y += dy[di[com]] * num
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
	fmt.Println(helper.Abs(x) + helper.Abs(y))
}
