package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"aoc"
)

func main() {
	var ops []string
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		terms := strings.Split(line, " ")
		if len(terms) < 2 {
			log.Print("INVALID LINE", line)
			continue
		}
		ops = append(ops, terms[0])
		nums = append(nums, aoc.ToInt(terms[1]))
	}

	result := 0
	tried := make([]bool, len(ops))
	for i := 0; i < len(ops); {
		if tried[i] {
			break
		}
		tried[i] = true

		if ops[i] == "acc" {
			result += nums[i]
			i++
		} else if ops[i] == "jmp" {
			i += nums[i]
		} else {
			// do nothing
			i++
		}

		if i < 0 {
			log.Print("Invalid Operation")
		}
	}

	fmt.Println(result)
}
