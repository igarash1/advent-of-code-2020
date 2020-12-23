package day23

import (
	"aoc"
	//"log"
	"strings"
)

func getInput(s string) []int {
	ss := strings.Split(s, "")
	cups := make([]int, len(ss))
	for i := range ss {
		cups[i] = aoc.ToInt(ss[i])
	}
	return cups
}

func getPos(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func toNum(nums []int) int {
	ret := 0
	s := getPos(nums, 1)
	for i := 1; i < len(nums); i++ {
		ret = ret*10 + nums[(s+i)%len(nums)]
	}
	return ret
}

func shiftLeft(nums []int, d int) []int {
	if d < 0 {
		d += len(nums)
	}
	newNums := make([]int, len(nums))
	for i := range nums {
		newNums[i] = nums[(i+d)%len(nums)]
	}
	return newNums
}

// O(#CUPS * #MOVES) (too slow for part 2)
func part1(input string) int {
	const MOVES = 100
	cups := getInput(input)
	lowest, highest := 1, 9
	cur := 0
	for t := 0; t < MOVES; t++ {
		cur = t % len(cups)
		pickUp := make(map[int]bool)
		for i := 1; i <= 3; i++ {
			pickUp[cups[(cur+i)%len(cups)]] = true
		}
		dest := cups[cur]
		for {
			dest--
			if dest < lowest {
				dest = highest
			}
			if !pickUp[dest] {
				break
			}
		}
		var newCups []int
		for _, cup := range cups {
			if pickUp[cup] {
				continue
			}
			newCups = append(newCups, cup)
			if cup == dest {
				for i := 1; i <= 3; i++ {
					newCups = append(newCups, cups[(cur+i)%len(cups)])
				}
			}
		}
		cups = shiftLeft(newCups, getPos(newCups, cups[cur])-cur)
	}
	return toNum(cups)
}

type Node struct {
	val  int
	prev *Node
	next *Node
}

type CircularList struct {
	head *Node
	tail *Node
}

func (list *CircularList) appendLast(t int) *Node {
	newNode := Node{t, list.tail, list.head}
	if list.head == nil {
		list.head = &newNode
		list.tail = &newNode
		newNode.prev = &newNode
		newNode.next = &newNode
	} else {
		list.tail.next = &newNode
		list.head.prev = &newNode
		list.tail = &newNode
	}
	return &newNode
}

func (list *CircularList) moveTo(from *Node, to *Node) {
	from.prev.next = from.next
	from.next.prev = from.prev
	from.prev = to
	from.next = to.next
	to.next.prev = from
	to.next = from
	if from == list.head {
		list.head = from.next
	}
	if to == list.tail {
		list.tail = from
	}
}

// O(#CUPS) in the initialization.
// O(#MOVES * constant) in the simulation.
func part2(input string) int {
	const NCUPS = 1000000
	const MOVES = 10000000
	firstCups := getInput(input)
	pos := make(map[int]*Node)

	var list CircularList
	for _, cup := range firstCups {
		pos[cup] = list.appendLast(cup)
	}
	for i := 10; i <= NCUPS; i++ {
		pos[i] = list.appendLast(i)
	}

	lowest, highest := 1, NCUPS
	cur := list.head
	for t := 0; t < MOVES; t++ {
		pCups := make([]*Node, 3)
		pick := cur.next
		for i := range pCups {
			pCups[i] = pick
			pick = pick.next
		}
		dest := cur.val
		for {
			dest--
			if dest < lowest {
				dest = highest
			}
			if pCups[0].val != dest && pCups[1].val != dest && pCups[2].val != dest {
				break
			}
		}
		last := pos[dest]
		for _, pCup := range pCups {
			list.moveTo(pCup, last)
			last = pCup
		}
		cur = cur.next
	}
	return pos[1].next.val * pos[1].next.next.val
}
