package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type INTERVAL struct {
	left, right int
}

func (ivl INTERVAL) Within(p int) bool {
	return ivl.left <= p && p <= ivl.right
}

func main() {
	ivlsMap := make(map[string][]INTERVAL)

	// scan the rules
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			log.Print("BLANK LINE")
			break
		}
		colonPos := strings.Index(line, ":")
		if colonPos == -1 {
			log.Print("INVALID RULE LINE", line)
		}
		fieldName := line[:colonPos]

		ss := bufio.NewScanner(strings.NewReader(line[colonPos+1:]))
		ss.Split(bufio.ScanWords)
		for ss.Scan() {
			sivl := ss.Text()
			if sivl == "or" {
				continue
			}
			var l, r int
			_, err := fmt.Sscanf(sivl, "%d-%d", &l, &r)
			if err != nil {
				log.Fatal(err)
			}
			ivlsMap[fieldName] = append(ivlsMap[fieldName], INTERVAL{l, r})
		}
	}

	// scan your ticket
	scanner.Scan()
	if scanner.Text() != "your ticket:" {
		log.Fatal("expected 'your ticket:', but ", scanner.Text())
	}

	scanner.Scan()
	yourTickets := splitToInts(scanner.Text())
	validYourTickets := getValidTickets(yourTickets, ivlsMap)

	// scan nearby tickets
	scanner.Scan()
	log.Print("BLANK LINE")
	if scanner.Scan(); scanner.Text() != "nearby tickets:" {
		log.Fatal("expected 'nearby tickets:', but ", scanner.Text())
	}

	var nearbyTickets [][]int
	for scanner.Scan() {
		nearbyTicket := splitToInts(scanner.Text())
		nearbyTickets = append(nearbyTickets, getValidTickets(nearbyTicket, ivlsMap))
	}

	// enumerate possible indices for each field
	fieldIdx := make(map[string][]int)
	for i, yt := range validYourTickets {
		for fName, ivls := range ivlsMap {
			if !isValidTicket(yt, ivls) {
				continue
			}
			possible := true
			for _, nearbyTicket := range nearbyTickets {
				if !isValidTicket(nearbyTicket[i], ivls) {
					possible = false
				}
			}
			if possible {
				log.Printf("the field of %d-th index can be %s", i, fName)
				fieldIdx[fName] = append(fieldIdx[fName], i)
			}
		}
	}

	// determine which field is which
	dFieldIdx := make(map[string]int)
	for len(dFieldIdx) < len(ivlsMap) {
		dIdx := -1
		for fName, idx := range fieldIdx {
			if len(idx) == 1 {
				dIdx = idx[0]
				log.Printf("the field of %d-th index is determined to be %s", dIdx, fName)
				dFieldIdx[fName] = dIdx
				break
			}
		}
		if dIdx == -1 {
			log.Fatal("something went wrong")
		}
		// remove the determined index from candidates
		for fName := range fieldIdx {
			fieldIdx[fName] = removeFromInts(fieldIdx[fName], dIdx)
		}
	}

	// compute the 'departure*' values on your ticket
	result := 1
	for fName, i := range dFieldIdx {
		if strings.HasPrefix(fName, "departure") {
			result *= yourTickets[i]
		}
	}
	fmt.Println(result)
}

func toInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
		return -1
	} else {
		return i
	}
}

func splitToInts(fields string) []int {
	var nums []int
	ss := strings.Split(fields, ",")
	for _, s := range ss {
		nums = append(nums, toInt(s))
	}
	return nums
}

// the discarded values are represented by -1
func getValidTickets(tickets []int, ivlsMap map[string][]INTERVAL) []int {
	var validTickets []int
	for _, t := range tickets {
		valid := false
		for _, ivls := range ivlsMap {
			if isValidTicket(t, ivls) {
				valid = true
				validTickets = append(validTickets, t)
				break
			}
		}
		if !valid {
			validTickets = append(validTickets, -1)
		}
	}
	return validTickets
}

func isValidTicket(ticket int, ivls []INTERVAL) bool {
	if ticket == -1 {
		return true
	}
	for _, ivl := range ivls {
		if ivl.Within(ticket) {
			return true
		}
	}
	return false
}

func removeFromInts(nums []int, target int) []int {
	var ret []int
	for _, v := range nums {
		if v != target {
			ret = append(ret, v)
		}
	}
	return ret
}
