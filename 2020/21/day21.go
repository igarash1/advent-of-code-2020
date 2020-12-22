package day21

import (
	"sort"

	//"fmt"
	//"os"
	"strings"
)

func getInput(input string) ([][]string, [][]string) {
	lines := strings.Split(input, "\n")

	var ingList [][]string
	var algList [][]string

	for _, line := range lines {
		if len(line) == 0 {
			//log.Print("BLANK LINE")
			continue
		}
		words := strings.Split(line, " ")
		var ings []string
		var algs []string
		inAlg := false
		for _, word := range words {
			name := strings.Trim(word, " (),")
			if name == "contains" {
				inAlg = true
			} else if inAlg {
				algs = append(algs, name)
			} else {
				ings = append(ings, name)
			}
		}
		ingList = append(ingList, ings)
		algList = append(algList, algs)
	}
	return ingList, algList
}

func getKeys(m map[string]string) []string {
	var uniqList []string
	for k, _ := range m {
		uniqList = append(uniqList, k)
	}
	return uniqList
}

func getUniqueList(ings [][]string, algs [][]string, exc map[string]bool) ([]string, []string) {
	ingSet := make(map[string]string)
	algSet := make(map[string]string)
	for i := range ings {
		for _, iName := range ings[i] {
			if exc[iName] {
				continue
			}
			ingSet[iName] = iName
		}
		for _, aName := range algs[i] {
			if exc[aName] {
				continue
			}
			algSet[aName] = aName
		}
	}
	ingUniqList := getKeys(ingSet)
	algUniqList := getKeys(algSet)
	return ingUniqList, algUniqList
}

func possible(graph map[string]map[string]bool, used map[string]bool, algs []string) bool {
	for _, aName := range algs {
		if used[aName] {
			continue
		}
		used[aName] = true
		for iName, _ := range graph[aName] {
			if used[iName] {
				continue
			}
			used[iName] = true
			if possible(graph, used, algs) {
				return true
			}
			used[iName] = false
		}
		return false
	}
	return true
}

// enumerate allergens that the target ingredient can contain
func enumerate(targetIng string, ings [][]string, algs [][]string, exclude map[string]bool) []string {
	// the amount of food
	N := len(algs)

	// assume given ingredients and allergens have unique names.
	ingUniqList, algUniqList := getUniqueList(ings, algs, exclude)
	graph := make(map[string]map[string]bool)
	// make the edges from allergens to ingredients
	for _, aName := range algUniqList {
		graph[aName] = make(map[string]bool)
		for _, iName := range ingUniqList {
			graph[aName][iName] = true
		}
	}

	// get the intersection of them
	for i := 0; i < N; i++ {
		for _, aName := range algs[i] {
			if exclude[aName] {
				continue
			}
			newEdges := make(map[string]bool)
			for _, iName := range ings[i] {
				if exclude[iName] {
					continue
				}
				if _, exist := graph[aName][iName]; exist {
					newEdges[iName] = true
				}
			}
			graph[aName] = newEdges
		}
	}

	var candAlgs []string
	for aName := range graph {
		for iName := range graph[aName] {
			if iName == targetIng {
				candAlgs = append(candAlgs, aName)
			}
		}
	}

	var possAlg []string
	for _, aName := range candAlgs {
		used := make(map[string]bool)
		used[targetIng] = true
		used[aName] = true
		if possible(graph, used, algUniqList) {
			possAlg = append(possAlg, aName)
		}
	}

	return possAlg
}

func part1(input string) int {
	ings, algs := getInput(input)
	poss := make(map[string]bool)
	result := 0
	for i := range ings {
		for _, iName := range ings[i] {
			if _, exist := poss[iName]; !exist {
				poss[iName] = len(enumerate(iName, ings, algs, make(map[string]bool))) > 0
			}
			if !poss[iName] {
				result++
			}
		}
	}

	return result
}

func part2(input string) string {
	ings, algs := getInput(input)
	exclude := make(map[string]bool)
	ingUniqList, _ := getUniqueList(ings, algs, exclude)

	// determine which ingredient contains which allergen
	ingToAlg := make(map[string]string)
	for found := true; found; {
		found = false
		for _, iName := range ingUniqList {
			if _, ok := ingToAlg[iName]; ok {
				continue
			}
			possAlg := enumerate(iName, ings, algs, exclude)
			if len(possAlg) == 1 {
				found = true
				aName := possAlg[0]
				ingToAlg[iName] = aName
				exclude[iName] = true
				exclude[aName] = true
			}
		}
	}

	algToIng := make(map[string]string)
	for iName, aName := range ingToAlg {
		algToIng[aName] = iName
	}
	algList := getKeys(algToIng)
	sort.Strings(algList)

	var result []string
	for _, aName := range algList {
		result = append(result, algToIng[aName])
	}
	return strings.Join(result, ",")
}
