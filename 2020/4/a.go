package main

import (
	"bufio"
	"fmt"
	// "log"
	"os"
	"strings"
)

func main() {
	eFlds := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fVlds := make(map[string]bool)
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
			if cidx > 0 && cidx < len(fld) - 1 {
				key := fld[:cidx]
				// value := fld[cidx+1:]
				fVlds[key] = true
			}
		}
		isEOF := !scanner.Scan()
		if isEOF || len(scanner.Text()) == 0 {
			valid := true
			for _, k := range eFlds {
				if !fVlds[k] {
					valid = false
					break
				}
			}
			if valid {
				// log.Println("valid")
				result++
			} else {
				// log.Println("invalid")
			}
			fVlds = make(map[string]bool)
			scanner.Scan()
		}

		if isEOF {
			break
		}
	}

	fmt.Println(result)
}
