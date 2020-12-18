package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Optype int

const (
	PLUS Optype = iota
	MULT
)

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func getRightPar(s string, i int) int {
	if s[i] != '(' {
		log.Fatal("NOT LEFT PARENTHESIS.")
	}
	cnt := 0
	for ; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		}
		if s[i] == ')' {
			cnt--
		}
		if cnt == 0 {
			return i
		}
	}
	log.Fatal("INVALID EXPRESSION: ", s)
	return -1
}

func calc(s string, l, r int) int {
	if s[r] == ' ' {
		return calc(s, l, r-1)
	}
	op := PLUS
	result := 0
	for ; l <= r; l++ {
		if s[l] == ' ' {
			continue
		}
		if s[l] == '+' {
			op = PLUS
		} else if s[l] == '*' {
			result *= calc(s, l+1, r)
			break
		} else {
			num := 0
			if s[l] == '(' {
				rp := getRightPar(s, l)
				num = calc(s, l+1, rp-1)
				l = rp
			} else {
				for ; l <= r && isDigit(s[l]); l++ {
					num = num*10 + int(s[l]-'0')
				}
				l--
			}
			if op == PLUS {
				result += num
			} else {
				result *= num
			}
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		result += calc(line, 0, len(line)-1)
	}
	fmt.Println(result)
}
