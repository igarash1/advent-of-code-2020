package main

import (
	"fmt"
	"log"
)

func main() {
	// O(N)

	result := 0
	for {
		var p1, p2 int
		var letter byte
		var pw string

		_, err := fmt.Scanf("%d-%d %c: %s", &p1, &p2, &letter, &pw)

		if err != nil {
			log.Print(err)
			break
		}

		cnt := 0
		if 0 < p1 && p1 <= len(pw) && byte(pw[p1-1]) == letter {
			cnt++
		}
		if 0 < p2 && p2 <= len(pw) && byte(pw[p2-1]) == letter {
			cnt++
		}

		if cnt == 1 {
			result++
		}
	}

	fmt.Println(result)
}
