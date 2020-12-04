package main

import (
	"fmt"
	"log"
)

func main() {
	// O(N)

	bucket := make([]bool, 2021)
	for {
		var expense int
		_, err := fmt.Scan(&expense)
		if err != nil {
			log.Print(err)
			break
		}
		if expense < 0 {
			log.Print("Negative expense.")
			// break
			continue
		}
		if expense > 2020 {
			continue
		}
		if bucket[2020-expense] {
			fmt.Println((2020 - expense) * expense)
		}
		bucket[expense] = true
	}
}
