package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSeatID(s string) int {
	if id, err := strconv.ParseInt(strings.Map(
		func(r rune) rune {
			switch {
			case r == 'B' || r == 'R':
				return '1'
			case r == 'F' || r == 'L':
				return '0'
			}
			return r
		}, s), 2, 32); err != nil {
		log.Fatal(err)
	} else {
		return int(id)
	}
	return -1 // don't happen
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
