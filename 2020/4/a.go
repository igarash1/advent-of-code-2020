package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isValid(v []bool) bool {
	for _, v := range v {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	eFlds := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fVlds := make([]bool, len(eFlds))
	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for {
		line := scanner.Text()
		input := bufio.NewScanner(strings.NewReader(line))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			fld := input.Text()
			if len(fld) == 0 {
				break
			}
			cidx := strings.Index(fld, ":")
			if cidx > 0 && cidx < len(fld) {
				id := fld[:cidx]
				for i, v := range eFlds {
					if id == v {
						fVlds[i] = true
						break
					}
				}
			}
		}
		isEOF := !scanner.Scan()
		if isEOF || len(scanner.Text()) == 0 {
			if isValid(fVlds) {
				log.Println("valid")
				result++
			} else {
				log.Println("invalid")
			}
			for i, _ := range fVlds {
				fVlds[i] = false
			}
			scanner.Scan()
		}

		if isEOF {
			break
		}
	}

	fmt.Println(result)
}
