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
		ans := make(map[byte]bool)
		for len(scanner.Text()) > 0 {
			line := scanner.Text()
			for _, c := range line {
				ans[byte(c)] = true
			}
			scanner.Scan()
		}
		result += len(ans)
	}

	fmt.Println(result)
}
