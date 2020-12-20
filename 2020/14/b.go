package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func setBit(x uint64, i int, b uint64) uint64 {
	return ((^(uint64(1) << i)) & x) | (b << i)
}

func writeMemory(mem map[uint64]uint64, mask string, addr uint64, value uint64) map[uint64]uint64 {
	possAddr := []uint64{0}
	for j, c := range mask {
		var nextPossAddr []uint64
		bitPos := 35 - j
		for _, pad := range possAddr {
			if c == 'X' {
				nextPossAddr = append(nextPossAddr, setBit(pad, bitPos, 0))
				nextPossAddr = append(nextPossAddr, setBit(pad, bitPos, 1))
			} else {
				pad = setBit(pad, bitPos, (addr>>(bitPos))&1)
				if c == '1' {
					pad = setBit(pad, bitPos, 1)
				}
				nextPossAddr = append(nextPossAddr, pad)
			}
		}
		possAddr = nextPossAddr
	}
	for _, a := range possAddr {
		mem[a] = value
	}
	return mem
}

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
