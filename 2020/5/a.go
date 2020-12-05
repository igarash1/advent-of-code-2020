package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
