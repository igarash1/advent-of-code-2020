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

func possible(graph map[string]map[string]bool, used map[string]bool, ings, algs []string) bool {
	for _, a := range algs {
		if used[a] {
			continue
		}
		used[a] = true
		for iName, _ := range graph[a] {
			if used[iName] {
				continue
			}
			//log.Println("Checking if", iName, "can contain", a)
			used[iName] = true
			if possible(graph, used, ings, algs) {
				//log.Println(iName, "contains", a)
				return true
			} else {
				//log.Println(iName, "cannot contain", a)
			}
			used[iName] = false
		}
		return false
	}
	return true
}

// enumerate allergens that the target ingredient can contain
func enumerate(targetIng string, ings [][]string, algs[][]string, exclude map[string]bool) []string {
	//log.Println("targetIng", targetIng)
	// making graph assuming they have unique names.
	graph := make(map[string]map[string]bool)
	for i := range ings {
		for _, iName := range ings[i] {
			if exclude[iName] {
				continue
			}
			if _, exist := graph[iName]; !exist {
				graph[iName] = make(map[string]bool)
			}
			for _, aName := range algs[i] {
				if exclude[aName] {
					continue
				}
				graph[iName][aName] = true
			}
		}
	}
	for i := range algs {
		for _, aName := range algs[i] {
			if exclude[aName] {
				continue
			}
			if _, exist := graph[aName]; !exist {
				graph[aName] = make(map[string]bool)
				for _, iName := range ings[i] {
					if exclude[iName] {
						continue
					}
					graph[aName][iName] = true
				}
			} else {
				newMap := make(map[string]bool)
				for _, iName := range ings[i] {
					if exclude[iName] {
						continue
					}
					if _, exist := graph[aName][iName]; exist {
						newMap[iName] = true
					}
				}
				graph[aName] = newMap
			}
		}
	}

	ingSetList, algSetList := getList(ings, algs, exclude)

	var possAlg []string
	for aName := range graph[targetIng] {
		if exclude[aName] || !graph[aName][targetIng] {
			continue
		}
		used := make(map[string]bool)
		used[targetIng] = true
		used[aName] = true
		// log.Println("Checking if", targetIng, "can contain", aName)
		if possible(graph, used, ingSetList, algSetList) {
			possAlg = append(possAlg, aName)
		}
	}

	return possAlg
}

func getList(ings [][]string, algs [][]string, exc map[string]bool) ([]string, []string) {
	ingSet := make(map[string]bool)
	algSet := make(map[string]bool)
	for i := range ings {
		for _, iName := range ings[i] {
			if exc[iName] {
				continue
			}
			ingSet[iName] = true
		}
		for _, aName := range algs[i] {
			if exc[aName] {
				continue
			}
			algSet[aName] = true
		}
	}
	var ingSetList []string
	for iName := range ingSet {
		ingSetList = append(ingSetList, iName)
	}
	var algSetList []string
	for aName := range algSet {
		algSetList = append(algSetList, aName)
	}
	return ingSetList, algSetList
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
				//log.Println(iName, "can't possibly contain any of the allergens.")
				result++
			}
		}
	}

	return result
}

func part2(input string) string {
	ings, algs := getInput(input)
	exc := make(map[string]bool)
	ingSetList, _ := getList(ings, algs, exc)

	ingToAlg := make(map[string]string)
	found := true
	for found {
		found = false
		for _, iName := range ingSetList {
			if exc[iName] {
				continue
			}
			if _, ok := ingToAlg[iName]; ok {
				continue
			}
			possAlg := enumerate(iName, ings, algs, exc)
			if len(possAlg) > 0 {
				//log.Println(iName, "can contain", possAlg)
			}
			if len(possAlg) == 1 {
				found = true
				aName := possAlg[0]
				ingToAlg[iName] = aName
				exc[iName] = true
				exc[aName] = true
			}
		}
	}

	algToIng := make(map[string]string)
	for iName, aName := range ingToAlg {
		//log.Println(iName, "contains", aName)
		algToIng[aName] = iName
	}
	var algList []string
	for aName := range algToIng {
		algList = append(algList, aName)
	}
	sort.Strings(algList)
	var result []string
	for _, aName := range algList {
		result = append(result, algToIng[aName])
	}
	return strings.Join(result, ",")
}
