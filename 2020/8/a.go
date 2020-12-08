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

func main() {
	ops := []string{}
	nums := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		terms := strings.Split(line, " ")
		if len(terms) < 2 {
			log.Print("INVALID LINE", line)
			continue
		}
		ops = append(ops, terms[0])
		nums = append(nums, toInt(terms[1]))
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
