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
	dp := make(map[int]int)
	dp[0] = 1
	for _, a := range adpts {
		dp[a] = dp[a-1]+dp[a-2]+dp[a-3]
	}
	fmt.Println(dp[adpts[len(adpts)-1]])
}
