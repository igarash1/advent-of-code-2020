package main

import (
	"aoc"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// returns the index of where the condition matched firstly, or -1 if not.
func match(rules map[int]string, s string, ruleNum int) int {
	if ruleStr := rules[ruleNum]; strings.Index(ruleStr, "\"") != -1 {
		// the rule is a single character rule
		//log.Println("single mathcing ", strings.Trim(ruleStr, "\""))
		p := strings.Index(s, strings.Trim(strings.TrimSpace(ruleStr), "\""))
		if p == -1 {
			return -1
		}
		return p
	} else {
		mRules := strings.Split(ruleStr, "|")
		//log.Println("Checking the one or multiple rules ", mRules)
		minIdx := len(s)
		for _, rule := range mRules {
			ruleNums := strings.Split(strings.TrimSpace(rule), " ")
			//log.Println("Checking the rule ", ruleNums)
			if len(ruleNums) == 0 {
				log.Fatal("something wrong", ruleNums)
			}
			matched := true
			matchedLast := -1
			for _, rs := range ruleNums {
				r := aoc.ToInt(rs)
				//log.Println("Checking if ",  s[matchedLast + 1:], " matches to the rule ", r)
				if p := match(rules, s[matchedLast+1:], r); p == -1 {
					//log.Print(s, " doesn't match to the rule ", r)
					matched = false
					break
				} else {
					matchedLast += p + 1
				}
			}
			if matched {
				//log.Println(s, " matched to ", rule)
				if matchedLast < minIdx {
					minIdx = matchedLast
				}
			}
		}

		if minIdx == len(s) {
			return -1
		} else {
			//log.Println("the minimum index of the instance of the rule", ruleNum, " is ", minIdx)
			return minIdx
		}
	}
}

func main() {
	rules := make(map[int]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			log.Println("BLANK LINE")
			break
		}
		ss := strings.Split(line, ":")
		ruleNum := aoc.ToInt(ss[0])
		rules[ruleNum] = ss[1]
	}

	result := 0
	for scanner.Scan() {
		msg := scanner.Text()
		if match(rules, msg, 0) == len(msg)-1 {
			log.Println(msg, " matched")
			result++
		}
	}

	fmt.Println(result)
}
