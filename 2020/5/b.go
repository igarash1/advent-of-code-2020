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
	scanner := bufio.NewScanner(os.Stdin)
	ids := [1 << 10]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 10 {
			// invalid line
			continue
		}
		id := getSeatID(line)
		ids[id] = true
	}

	for n := 1; n < (1<<10)-1; n++ {
		if ids[n-1] && !ids[n] && ids[n+1] {
			fmt.Println(n)
		}
	}
}
