package day23

import (
	"aoc"
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

// O(#CUPS * #MOVES) (too slow for part 2. see part2() for better solution)
func part1(input string) int {
	const MOVES = 100
	cups := getInput(input)
	lowest, highest := 1, 9
	NCUPS := len(cups)
	cur := 0
	for t := 0; t < MOVES; t++ {
		cur = t % len(cups)
		pickUp := make(map[int]bool)
		for i := 1; i <= 3; i++ {
			pickUp[cups[(cur+i)%NCUPS]] = true
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
					newCups = append(newCups, cups[(cur+i)%NCUPS])
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

func (list *CircularList) move(from *Node, to *Node) {
	if from == list.head {
		list.head = from.next
	}
	from.prev.next = from.next
	from.next.prev = from.prev
	from.prev = to
	from.next = to.next
	to.next.prev = from
	to.next = from
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
	pos := make([]*Node, NCUPS+1)

	var list CircularList
	for _, cup := range firstCups {
		pos[cup] = list.appendLast(cup)
	}
	for i := 10; i <= NCUPS; i++ {
		pos[i] = list.appendLast(i)
	}

	cur := list.head
	for mv := 0; mv < MOVES; mv++ {
		p1 := cur.next
		p2 := cur.next.next
		p3 := cur.next.next.next
		dest := cur.val
		for {
			dest = (dest-2+NCUPS)%NCUPS + 1
			if dest != p1.val && dest != p2.val && dest != p3.val {
				break
			}
		}
		list.move(p1, pos[dest])
		list.move(p2, p1)
		list.move(p3, p2)
		cur = cur.next
	}

	return pos[1].next.val * pos[1].next.next.val
}
