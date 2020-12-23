package aoc

import (
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

func ToInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		log.Fatal(err)
		return -1
	} else {
		return i
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetMin(nums []int) int {
	minv := nums[0]
	for _, x := range nums {
		minv = IntMin(minv, x)
	}
	return minv
}

func GetMax(nums []int) int {
	maxv := nums[0]
	for _, x := range nums {
		maxv = IntMax(maxv, x)
	}
	return maxv
}

func ReadStringFromFile(fileName string) string {
	inputs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(inputs)
}

func ReadString(r io.Reader) string {
	inputs, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(inputs)
}
