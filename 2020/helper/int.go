package helper

import (
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
