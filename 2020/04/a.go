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
	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fVlds := make(map[string]bool)
		for len(scanner.Text()) > 0 {
			line := scanner.Text()
			input := bufio.NewScanner(strings.NewReader(line))
			input.Split(bufio.ScanWords)
			for input.Scan() {
				fld := input.Text()
				if len(fld) == 0 {
					break
				}
				cidx := strings.Index(fld, ":")
				if cidx > 0 && cidx < len(fld)-1 {
					key := fld[:cidx]
					// value := fld[cidx+1:]
					fVlds[key] = true
				}
			}
			scanner.Scan()
		}
		valid := true
		for _, k := range eFlds {
			if !fVlds[k] {
				valid = false
			}
		}
		if valid {
			// log.Println("valid")
			result++
		} else {
			// log.Println("invalid")
		}
	}

	fmt.Println(result)
}
