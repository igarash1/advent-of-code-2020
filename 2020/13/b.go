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

/*
* mode_inverse(x, m):
* compute the inverse a of x mod m,
* solving ax+my=1 by extended-gcd.
 */
func mod_inverse(x, m int) int {
	if m == 1 {
		return 0
	}

	M := m
	x0, x1 := 0, 1
	for x > 1 {
		q := x / m
		m, x = x%m, m
		x0, x1 = x1-q*x0, x0
	}

	return (x1 + M) % M
}

func main() {
	var ts int
	busRem := []int{}
	busIDs := []int{}

	fmt.Scan(&ts)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	raw_input := strings.Replace(scanner.Text(), " ", "", -1)
	input := strings.Split(raw_input, ",")
	for i, s := range input {
		if s != "x" {
			bid := toInt(s)
			busRem = append(busRem, (bid-(i%bid))%bid)
			busIDs = append(busIDs, bid)
		}
	}

	// we assume bus IDs are coprime.
	P := 1
	for _, b := range busIDs {
		P *= b
	}

	/*
	* we have to find the solution x s.t.
	* 	busRem[i] = result (mod busIDs[i]) for all i.
	* By Chinese Remainder Theorem, we know there is a unique solution in
	* mod LCM(busIDs) and it is smallest.
	* so we just only have to find a solution and take mod.
	* let's define P[i] and P[i] ^ -1:
	* 	 P[i] := LCM(busIDs) / busIDs[i],
	* 	 P[i]^-1 := the inverse of P[i] under the modulo busIDs[i].
	*  note that P[i] and busIDs[i] are coprime. and then,
	* 	busRem[0] * P[0]^-1 * P[0] + ... + busRem[n-1] * P[n-1]^-1 * P[n-1],
	* where
	* is one solution satysfing the mod equations since
	* 	busRem[j] * P[j]^-1 * P[j] = 0 (mod P[i]) (for all j st. i !=
	* 	j).
	* note that P[j] is a mutiple of P[i].
	 */
	result := 0
	for i, id := range busIDs {
		pi := P / id
		result += busRem[i] * mod_inverse(pi, id) * pi
	}

	fmt.Println(result % P)
}
