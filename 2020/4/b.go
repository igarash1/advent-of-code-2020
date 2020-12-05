package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
		// log.Fatal("ERROR in convertin string to int ", err)
	}
	return v
}

func validValueRange(key string, value string) bool {
	var num int
	if key == "hgt" {
		if len(value) < 3 {
			return false
		}
		num = toInt(value[:len(value)-2])
	} else {
		num = toInt(value)
	}

	if key == "byr" {
		return 1920 <= num && num <= 2002
	}
	if key == "iyr" {
		return 2010 <= num && num <= 2020
	}
	if key == "eyr" {
		return 2020 <= num && num <= 2030
	}
	if key == "hgt" {
		suffix := value[len(value)-2:]
		if suffix == "cm" {
			return 150 <= num && num <= 193
		}
		if suffix == "in" {
			return 59 <= num && num <= 76
		}
		return false
	}

	return true
}

func main() {
	eFlds := map[string]string{
		"byr": `^[0-9]{4}$`,
		"iyr": `^[0-9]{4}$`,
		"eyr": `^[0-9]{4}$`,
		"hgt": `^([0-9]{3}cm|[0-9]{2}in)$`,
		"hcl": `^#[0-9a-f]{6}$`,
		"ecl": `^(amb|blu|brn|gry|grn|hzl|oth)$`,
		"pid": `^[0-9]{9}$`}

	rFlds := make(map[string]*regexp.Regexp)
	for k, _ := range eFlds {
		r, err := regexp.Compile(eFlds[k])
		if err != nil {
			log.Fatal("regex error ", err)
		}
		rFlds[k] = r
	}

	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fVlds := make(map[string]string)
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
				if cidx > 0 && cidx < len(fld) {
					id := fld[:cidx]
					if id != "cid" {
						fVlds[id] = fld[cidx+1:]
					}
				}
			}
			scanner.Scan()
		}

		valid := true
		for k, regx := range rFlds {
			if regx.MatchString(fVlds[k]) {
				if !validValueRange(k, fVlds[k]) {
					valid = false
				}
			} else {
				valid = false
			}
		}

		if len(fVlds) > 0 && valid {
			result++
		}
	}

	fmt.Println(result)
}
