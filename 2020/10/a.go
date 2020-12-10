package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	adpts := []int{}
	for {
		var x int
		if _, err := fmt.Scan(&x); err != nil {
			log.Print(err)
			break
		}
		adpts = append(adpts, x)
	}
	sort.Slice(adpts, func(i, j int) bool { return adpts[i] < adpts[j] })
	dj1, dj3, curj := 0, 1, 0
	for _, a := range adpts {
		if d := a - curj; d == 1 {
			dj1++
		} else if d == 3 {
			dj3++
		}
		curj = a
	}
	fmt.Println(dj1 * dj3)
}
