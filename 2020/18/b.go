package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	LL(1) Parsing.
	BNF:
		expr	::= expr '*' sum			| sum
		sum		::= sum '+' term			| term
		term	::= (expr)				| Number
	BNF (after removal of left recursion):
		expr	::= sum ('*' sum)*		| sum
		sum		::= sum ('+' sum)*		| term
		term	::= (expr)				| Number
*/

type Number int
type Index int

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func expr(s string, i Index) (Number, Index) {
	var result Number
	result, i = sum(s, i)
	for s[i] == '*' {
		rTerm, j := sum(s, i+1)
		result *= rTerm
		i = j
	}
	return result, i
}

func sum(s string, i Index) (Number, Index) {
	var result Number
	result, i = term(s, i)
	for s[i] == '+' {
		rTerm, j := term(s, i+1)
		result += rTerm
		i = j
	}
	return result, i
}

func term(s string, i Index) (Number, Index) {
	if s[i] == '(' {
		result, i := expr(s, i+1)
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
