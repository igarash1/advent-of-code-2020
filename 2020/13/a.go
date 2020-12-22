package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"aoc"
)

func main() {
	var ts int
	var busIds []int

	_, _ = fmt.Scan(&ts)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			raw_input := strings.Replace(scanner.Text(), " ", "", -1)
			input := strings.Split(raw_input, ",")
			for _, s := range input {
				if s != "x" {
					busIds = append(busIds, aoc.ToInt(s))
				}
			}
		}
	}

	best := -1
	result := -1
	for _, b := range busIds {
		if ts%b == 0 {
			best = 0
			result = 0
		} else {
			cand := (ts/b)*b + b
			log.Print(cand)
			if best == -1 || cand < best {
				best = cand
				result = (cand - ts) * b
			}
		}
	}

	fmt.Println(result)
}
