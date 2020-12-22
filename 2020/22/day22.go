package day22

import (
	"aoc"
	"strings"
)

func getDeck(input string) ([]int, []int) {
	players := strings.Split(input, "\n\n")
	var deck1 []int
	lines := strings.Split(players[0], "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line == "Player 1:" {
			continue
		}
		deck1 = append(deck1, aoc.ToInt(line))
	}
	var deck2 []int
	lines = strings.Split(players[1], "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line == "Player 2:" {
			continue
		}
		deck2 = append(deck2, aoc.ToInt(line))
	}
	return deck1, deck2
}

func getScore(deck []int) int {
	score := 0
	for i, v := range deck {
		score += (len(deck) - i) * v
	}
	return score
}

func regularCombat(deck1 []int, deck2 []int) int {
	for len(deck1) > 0 && len(deck2) > 0 {
		a, b := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		if a > b {
			deck1 = append(deck1, a, b)
		} else {
			deck2 = append(deck2, b, a)
		}
	}
	return getScore(deck1) + getScore(deck2)
}

func part1(input string) int {
	deck1, deck2 := getDeck(input)
	return aoc.Abs(regularCombat(deck1, deck2))
}

const B1 uint64 = 5
const B2 uint64 = 7

func getHash(deck1 []int, deck2 []int) uint64 {
	var hash uint64 = 0
	for _, v := range deck1 {
		hash = hash*B1 + uint64(v)
	}
	for _, v := range deck2 {
		hash = hash*B2 + uint64(v)
	}
	return hash
}

func copyInts(src []int) []int {
	dst := make([]int, len(src))
	for i := range src {
		dst[i] = src[i]
	}
	return dst
}

func recursiveCombat(deck1, deck2 []int) int {
	prev := make(map[uint64]bool)
	for r := 1; len(deck1) > 0 && len(deck2) > 0; r++ {
		hash := getHash(deck1, deck2)
		a, b := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		if prev[hash] {
			deck1 = append(deck1, a, b)
			continue
		} else if a <= len(deck1) && b <= len(deck2) {
			if recursiveCombat(copyInts(deck1[:a]), copyInts(deck2[:b])) > 0 {
				deck1 = append(deck1, a, b)
			} else {
				deck2 = append(deck2, b, a)
			}
		} else {
			if a > b {
				deck1 = append(deck1, a, b)
			} else {
				deck2 = append(deck2, b, a)
			}
		}
		prev[hash] = true
	}

	// return positive score if player 1 wins, otherwise negative score
	if len(deck1) > 0 {
		return getScore(deck1)
	}
	return -getScore(deck2)
}

func part2(input string) int {
	deck1, deck2 := getDeck(input)
	return aoc.Abs(recursiveCombat(deck1, deck2))
}
