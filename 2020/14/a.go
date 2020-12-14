package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
)

func main() {
	var curMask string
	mem := make(map[uint64]uint64)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:4] == "mask" {
			_, err := fmt.Sscanf(line, "mask = %s", &curMask)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			var addr uint64
			var value uint64
			_, err := fmt.Sscanf(line, "mem[%d] = %d", &addr, &value)
			if err != nil {
				log.Fatal(err)
			}
			mem = writeMemory(mem, curMask, addr, value)
		}
	}

	var result uint64 = 0
	for _, v := range mem {
		result += v
	}
	fmt.Println(result)
}

func setBit(x uint64, i int, b uint64) uint64 {
	return ((^(uint64(1) << i)) & x) | (b << i)
}

func writeMemory(mem map[uint64]uint64, mask string, addr uint64, value uint64) map[uint64]uint64 {
	for j, c := range mask {
		bitPos := 35 - j
		if c != 'X' {
			value = setBit(value, bitPos, uint64(c-'0'))
		}
	}
	mem[addr] = value
	return mem
}
