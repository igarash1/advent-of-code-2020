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
	if i, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
		return -1
	} else {
		return i
	}
}

func createRegex(rules map[int]string, ruleNum int) string {
	// fix the rule 8 and 11
	if ruleNum == 8 {
		return "(" + createRegex(rules, 42) + "+)"
	}
	if ruleNum == 11 {
		// simply make the nested regex for the recusive patterns: 42 31 | 42 42 31 31 |  ...
		const MAX_DEPTH = 5 // the maximum depth of the nest, I guess roughly from the data
		regexResult := ""
		cur := ""
		rule42s := createRegex(rules, 42)
		rule31s := createRegex(rules, 31)
		for d := 0;d < MAX_DEPTH;d++ {
			cur = "(" + rule42s + cur + rule31s + ")"
			if len(regexResult) > 0 {
				regexResult += "|"
			}
			regexResult += cur
		}
		return "(" + regexResult + ")"
	}

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
				r := toInt(rs)
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
		//log.Println(ss[0], "and", ss[1])
		ruleNum := toInt(ss[0])
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
			log.Println(msg, " matched to the updated rule 0")
			result++
		}
	}

	fmt.Println(result)
}
