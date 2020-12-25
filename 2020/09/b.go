package main

import (
	"fmt"
	"log"
)

func main() {
	nums := []int{}
	sums := []int{}
	occ := make(map[int]int)
	sum := 0
	for i := 0; ; i++ {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			log.Print(err)
			break
		}
		sum += num
		sums = append(sums, sum)
		nums = append(nums, num)
		if i < 25 {
			occ[num]++
		}
	}

	pos := -1
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
			pos = i
			break
		}

		occ[nums[i]]++
		occ[nums[i-25]]--
	}

	for i := 0; i < pos; i++ {
		minv, maxv := nums[i], nums[i]
		for j := i + 1; j < pos; j++ {
			if nums[j] > maxv {
				maxv = nums[j]
			}
			if nums[j] < minv {
				minv = nums[j]
			}
			if sums[j]-sums[i]+nums[i] == nums[pos] {
				fmt.Println(minv + maxv)
				return
			}
		}
	}
}
