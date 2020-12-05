package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSeatID(s string) int {
	id := 0
	for _, c := range s {
		id <<= 1
		if c == 'B' || c == 'R' {
			id += 1
		}
	}
	return id
}

func main() {
	result := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 10 {
			// invalid line
			continue
		}
		id := getSeatID(line)
		if id > result {
			result = id
		}
	}

	fmt.Println(result)
}
