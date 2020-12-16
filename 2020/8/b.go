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

func execute(ops []string, nums []int) (int, bool) {
	result := 0
	tried := make([]bool, len(ops))
	i := 0
	for i < len(ops) {
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
			break
		}
	}

	return result, i == len(ops)
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

	swap := func(i int) {
		if ops[i] == "jmp" {
			ops[i] = "nop"
		} else if ops[i] == "nop" {
			ops[i] = "jmp"
		}
	}

	for i := range ops {
		if ops[i] == "acc" {
			continue
		}
		swap(i)
		result, success := execute(ops, nums)
		if success {
			fmt.Println(result)
			break
		}
		swap(i)
	}

}
