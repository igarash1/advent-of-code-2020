package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func validValue(key string, value string) bool {
	kMatched, err := regexp.Match(`byr|iyr|eyr|hgt`, []byte(key))
	if err != nil {
		log.Fatal("regex error ", err)
		return false
	}
	if !kMatched {
		// log.Print("need no value validation ", key)
		return true
	}

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

	return false // don't happen
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

	aVlds := map[string]map[string]string{}
	for k, _ := range eFlds {
		aVlds[k] = make(map[string]string)
	}
	fVlds := make(map[string]string)
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
				if id != "cid" {
					fVlds[id] = fld[cidx+1:]
				}
			}
		}

		isEOF := !scanner.Scan()
		if isEOF || len(scanner.Text()) == 0 {
			valid := true
			for k, rs := range eFlds {
				r, err := regexp.Compile(rs)
				if err != nil {
					log.Fatal("regex error ", err)
				}
				if r.MatchString(fVlds[k]) {
					if !validValue(k, fVlds[k]) {
						valid = false
					}
				} else {
					valid = false
				}
			}

			if valid {
				for k, v := range fVlds {
					aVlds[k][v] = v
				}
				log.Println("VALID ", fVlds)
				result++
			}
			fVlds = make(map[string]string)
			scanner.Scan()
		}
		// log.Print("NEXT LINE ", scanner.Text())

		if isEOF {
			break
		}
	}

	log.Print("APPEARED VALUES")
	for k, v := range aVlds {
		log.Print(k, ":")
		keys := []string{}
		for k, _ := range v {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		log.Print(keys)
	}

	fmt.Println(result)
}
