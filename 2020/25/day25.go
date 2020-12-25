package day25

import (
	"aoc"
	"strings"
)

const MOD = 20201227

func transform(loopSize, subjectKey int) int {
	v := 1
	for i := 0; i < loopSize; i++ {
		v *= subjectKey
		v %= MOD
	}
	return v
}

func detransform(subjectKey, expected int) int {
	v := 1
	loopSize := 0
	for v != expected {
		v *= subjectKey
		v %= MOD
		loopSize++
	}
	return loopSize
}

func handshake(cardLoopSize, doorLoopSize, cardPubKey, doorPubKey int) (int, bool) {
	cardEncryptKey := transform(cardLoopSize, doorPubKey)
	doorEncryptKey := transform(doorLoopSize, cardPubKey)
	return cardEncryptKey, cardEncryptKey == doorEncryptKey
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	cardPubKey := aoc.ToInt(lines[0])
	doorPubKey := aoc.ToInt(lines[1])
	cardLoopSize := detransform(7, cardPubKey)
	doorLoopSize := detransform(7, doorPubKey)
	encryptKey, _ := handshake(cardLoopSize, doorLoopSize, cardPubKey, doorPubKey)
	return encryptKey
}
