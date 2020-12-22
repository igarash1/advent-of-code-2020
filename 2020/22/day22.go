package day22

import (
	"aoc"
	"strings"
)

// returns the signed score of the winner. + if player 1 wins, - if player 2 wins.
func getDecksFromInput(input string) ([]int, []int) {
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

// returns the score of the winner.
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
	deck1, deck2 := getDecksFromInput(input)
	return aoc.Abs(regularCombat(deck1, deck2))
}

// the bases for hashing, use primes to avoid hash collision
const B1 uint64 = 5
const B2 uint64 = 7

// use hash to represent a state of the game while avoiding hash collision.
// make sure it return the same result if the given state of decks is the same.
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

// returns the signed score of the winner. + if player 1 wins, - if player 2 wins.
func recursiveCombat(deck1, deck2 []int) int {
	prev := make(map[uint64]bool)
	for r := 1; len(deck1) > 0 && len(deck2) > 0; r++ {
		hash := getHash(deck1, deck2)
		a, b := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		if prev[hash] {
			// there was a previous round with the same decks
			deck1 = append(deck1, a, b)
			continue
		} else if a <= len(deck1) && b <= len(deck2) {
			// sub-game
			if recursiveCombat(copyInts(deck1[:a]), copyInts(deck2[:b])) > 0 {
				deck1 = append(deck1, a, b)
			} else {
				deck2 = append(deck2, b, a)
			}
		} else {
			// just play a regular game
			if a > b {
				deck1 = append(deck1, a, b)
			} else {
				deck2 = append(deck2, b, a)
			}
		}
		prev[hash] = true
	}

	// return positive score if player 1 wins, otherwise negative score
	return getScore(deck1) - getScore(deck2)
}

func part2(input string) int {
	deck1, deck2 := getDecksFromInput(input)
	return aoc.Abs(recursiveCombat(deck1, deck2))
}
