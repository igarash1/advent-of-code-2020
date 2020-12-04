package main

import (
	"fmt"
	"log"
)

func main() {
	// O(N * L), where L is the maximum length of the password.

	result := 0
	for {
		var lo, hi int
		var letter byte
		var pw string

		_, err := fmt.Scanf("%d-%d %c: %s", &lo, &hi, &letter, &pw)

		if err != nil {
			log.Print(err)
			break
		}

		cnt := 0
		for _, c := range pw {
			if byte(c) == letter {
				cnt++
			}
		}
		if lo <= cnt && cnt <= hi {
			result++
		}
	}

	fmt.Println(result)
}
