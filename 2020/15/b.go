package main

import (
	"fmt"
	"log"
)

func main() {
	// use array for performance
	const TURN = 30000000
	var prev [TURN]int
	turn := 1
	last := -1
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			log.Print(err)
			break
		}
		last = num
		prev[num] = turn
		turn++
	}

	for ;turn <= TURN;turn++ {
		var next int
		if prev[last] > 0 {
			next = turn - prev[last] - 1
		} else {
			next = 0
		}
		prev[last] = turn - 1
		// log.Printf("%d-th number is %d\n", turn, next)
		last = next
	}

	fmt.Println(last)
}
