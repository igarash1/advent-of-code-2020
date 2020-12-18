package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Number int
type Index int

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func expr(s string, i Index) (Number, Index){
	var result Number
	result, i = term(s, i)
	for s[i] == '*' || s[i] == '+' {
		rterm, j := term(s, i + 1)
		if s[i] == '*' {
			result *= rterm
		} else {
			result += rterm
		}
		i = j
	}
	return result, i
}

func term(s string, i Index) (Number, Index) {
	if s[i] == '(' {
		result, i := expr(s, i + 1)
		return result, i + 1 // skip for right parenthesis
	} else {
		return number(s, i)
	}
}

func number(s string, i Index) (Number, Index) {
	var num Number = 0
	for ; isDigit(s[i]); i++ {
		num = num*10 + Number(s[i]-'0')
	}
	return num, i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var result Number = 0
	for scanner.Scan() {
		line := scanner.Text()
		// remove all whitespaces and padding for simpler implementation
		line = strings.ReplaceAll(line, " ", "") + " "
		v, _ := expr(line, 0)
		log.Println(line, " = ", v)
		result += v
	}
	fmt.Println(result)
}
