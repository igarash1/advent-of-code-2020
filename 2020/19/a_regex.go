package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"aoc"
)

func createRegex(rules map[int]string, ruleNum int) string {
	if ruleStr, ok := rules[ruleNum]; !ok {
		log.Fatal("something wrong")
		return ""
	} else if strings.Index(ruleStr, "\"") != -1 {
		// the rule is a single character rule
		return strings.Trim(strings.TrimSpace(ruleStr), "\"")
	} else {
		mRules := strings.Split(ruleStr, "|")
		regex := ""
		for i, rule := range mRules {
			if i > 0 {
				regex += "|"
			}
			regex += "("
			ruleNums := strings.Split(strings.TrimSpace(rule), " ")
			if len(ruleNums) == 0 {
				log.Fatal("something wrong", ruleNums)
			}
			for _, rs := range ruleNums {
				r := aoc.ToInt(rs)
				regex += createRegex(rules, r)
			}
			regex += ")"
		}
		if len(mRules) > 1 {
			regex = "(" + regex + ")"
		}
		return regex
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

	ruleZeroRegex, err := regexp.Compile("^" + createRegex(rules, 0) + "$")
	if err != nil {
		log.Fatal("regex error ", err)
	}

	result := 0
	for scanner.Scan() {
		msg := scanner.Text()
		if ruleZeroRegex.MatchString(msg) {
			log.Println(msg, " matched to the rule 0")
			result++
		}
	}

	fmt.Println(result)
}
