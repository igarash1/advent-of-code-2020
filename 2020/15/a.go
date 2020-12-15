package main

import (
	"fmt"
	"log"
)

func main() {
	prev := make(map[int]int)
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
	
	for ;turn <= 2020;turn++ {
		var next int
		if _, ok := prev[last]; ok {
			next = turn - prev[last] - 1
		} else {
			next = 0
		}
		prev[last] = turn - 1
		log.Printf("%d-th number is %d\n", turn, next)
		last = next
	}

	fmt.Println(last)
}
