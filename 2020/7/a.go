package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func trim(s string) string {
	return strings.Trim(s, " 1234567890")
}

func getBagName(s string) string {
	return trim(s[:strings.Index(s, "bag")])
}

func main() {
	graph := make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		bags := strings.Split(line, "contain")
		log.Print(bags)
		if len(bags) < 2 {
			log.Print("INVALID LINE", line)
			continue
		}

		lBagName := getBagName(bags[0])
		rBags := strings.Split(bags[1], ",")
		for _, rBag := range rBags {
			rBagName := getBagName(rBag)
			if rBagName == "no other" {
				continue
			}
			graph[rBagName] = append(graph[rBagName], lBagName)
		}
	}

	visited := make(map[string]bool)
	que := []string{"shiny gold"}
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		if visited[u] {
			continue
		}
		visited[u] = true
		for _, v := range graph[u] {
			que = append(que, v)
		}
	}

	fmt.Println(len(visited) - 1)
}
