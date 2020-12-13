package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func mod_inverse(x, m int) int{
	if m == 1 {
		return 0
	}

	m0 := m
	x0, x1 := 0, 1
	for x > 1 {
		q := x / m
		t := m
		m, x = x % m, t
		t = x0
		x0 = x1 - q * x0;
		x1 = t;
	}

	if x1 < 0 {
		x1 += m0;
	}

	return x1;
}

func gcd(a, b int) int {
	if b < a {
		return gcd(b, a)
	}
	if a == 0 {
		return b
	}
	return gcd(b % a, a)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	var ts int
	busRem := []int{}
	busIds := []int{}

	fmt.Scan(&ts)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	raw_input := strings.Replace(scanner.Text(), " ", "", -1)
	input := strings.Split(raw_input, ",")
	for i, s := range input {
		if s != "x" {
			bid := toInt(s)
			busRem = append(busRem, (bid - (i % bid)) % bid)
			busIds = append(busIds, bid)
		}
	}

	P, LCM := 1, 1
	for _, b := range busIds {
		P *= b
		LCM = lcm(LCM, b)
	}

	result := 0
	for i, bid := range busIds {
		pp := LCM / bid
		result += busRem[i] * mod_inverse(pp, bid) * pp
	}

	fmt.Println(result % LCM)
}

