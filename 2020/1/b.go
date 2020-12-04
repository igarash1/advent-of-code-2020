package main

import (
	"fmt"
	"log"
)

func main() {
	// O(N^2)

	var exp []int
	bucket := make([]bool, 2021)
	for {
		var x int
		_, err := fmt.Scan(&x)
		if err != nil {
			log.Print(err)
			break
		}
		if x < 0 || x > 2020 {
			log.Print("Invalid input:", x)
			// break
			continue
		}
		exp = append(exp, x)
		bucket[x] = true
	}

	for i, _ := range exp {
		for j := i + 1; j < len(exp); j++ {
			if exp[i]+exp[j] > 2020 {
				continue
			}
			if bucket[2020-exp[i]-exp[j]] {
				fmt.Println(exp[i] * exp[j] * (2020 - exp[i] - exp[j]))
			}
		}
	}
}
