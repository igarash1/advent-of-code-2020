package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cnt := 0
		tot := make(map[byte]int)
		for len(scanner.Text()) > 0 {
			cnt++
			ans := make(map[byte]bool)
			line := scanner.Text()
			for _, c := range line {
				ans[byte(c)] = true
			}
			for k, _ := range ans {
				tot[k]++
			}
			scanner.Scan()
		}
		for _, v := range tot {
			if v == cnt {
				result++
			}
		}
	}

	fmt.Println(result)
}
