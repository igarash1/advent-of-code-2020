package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"../helper"
)

// returns the index of where the condition matched firstly, or -1 if not.
func match(rules map[int]string, s string, ruleNum int) int {
	if len(s) == 0 {
		return -1
	}
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
		maxIdx := len(s)
		for _, rule := range mRules {
			ruleNums := strings.Split(strings.TrimSpace(rule), " ")
			//log.Println("Checking the rule ", ruleNums)
			if len(ruleNums) == 0 {
				log.Fatal("something wrong", ruleNums)
			}
			matched := true
			matchedLast := -1
			for _, rs := range ruleNums {
				r := helper.ToInt(rs)
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
				if matchedLast > maxIdx {
					maxIdx = matchedLast
				}
			}
		}

		if maxIdx == len(s) {
			return -1
		} else {
			//log.Println("the minimum index of the instance of the rule", ruleNum, " is ", minIdx)
			return maxIdx
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
		ruleNum := helper.ToInt(ss[0])
		rules[ruleNum] = ss[1]
	}

	// fix the rule 8 and 11
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"

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
