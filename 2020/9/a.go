package main

import (
	"fmt"
	"log"
)

func main() {
	nums := []int{}
	occ := make(map[int]int)
	for i := 0; ; i++ {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			log.Print(err)
			break
		}
		nums = append(nums, num)
		if i < 25 {
			occ[num]++
		}
	}

	for i := 25; i < len(nums); i++ {
		valid := false
		for j := i - 25; j < i; j++ {
			if occ[nums[i]-nums[j]] > 0 {
				log.Printf("%d = %d + %d\n", nums[i], nums[i]-nums[j], nums[j])
				valid = true
				break
			}
		}

		if !valid {
			fmt.Println(nums[i])
			break
		}

		occ[nums[i]]++
		occ[nums[i-25]]--
	}
}
