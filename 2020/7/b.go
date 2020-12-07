package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	num  int
	name string
}

func getNum(s string) int {
	s = strings.Trim(s, " ")
	if i, err := strconv.Atoi(s[:strings.Index(s, " ")]); err != nil {
		log.Fatal(s, err)
		return -1
	} else {
		return i
	}
}

func trim(s string) string {
	return strings.Trim(s, " 1234567890")
}

func getBagName(s string) string {
	return trim(s[:strings.Index(s, "bag")])
}

func main() {
	graph := make(map[string][]Node)

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
			node := Node{getNum(rBag), rBagName}
			log.Print(node)
			graph[lBagName] = append(graph[lBagName], node)
		}
	}

	que := []Node{Node{1, "shiny gold"}}
	result := 0
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		result += u.num
		for _, v := range graph[u.name] {
			que = append(que, Node{u.num * v.num, v.name})
		}
	}

	fmt.Println(result - 1)
}
