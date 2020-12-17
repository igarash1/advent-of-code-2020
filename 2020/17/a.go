package main

import (
	"bufio"
	"fmt"
	"os"
)

type State int

const (
	ACTIVE State = iota
	INACTIVE
)

func getState(c byte) State {
	if c == '#' {
		return ACTIVE
	} else {
		return INACTIVE
	}
}

const DIM = 3

type Point struct {
	crd [DIM]int
}

func newPoint(x, y int) Point {
	return Point{[DIM]int{x, y}}
}

func getCopy(source map[Point]State) map[Point]State {
	target := make(map[Point]State, len(source))
	for k, v := range source {
		target[k] = v
	}
	return target
}

func getCopyWithInit(source map[Point]State, initValue State) map[Point]State {
	target := make(map[Point]State, len(source))
	for k := range source {
		target[k] = initValue
	}
	return target
}

// enumerate neighbors of a point including itself
func getNeighbors(p Point) []Point {
	res := []Point{p}
	for d := 0; d < DIM; d++ {
		var next []Point
		for _, cur := range res {
			for r := -1; r <= 1; r++ {
				np := cur
				np.crd[d] += r
				next = append(next, np)
			}
		}
		res = next
	}
	return res
}

func extendCube(cubes map[Point]State) map[Point]State {
	for p, s := range cubes {
		if s == INACTIVE {
			// do not extend if unnecessary
			continue
		}
		neighs := getNeighbors(p)
		for _, np := range neighs {
			if _, ok := cubes[np]; !ok {
				cubes[np] = INACTIVE
			}
		}
	}
	return cubes
}

func countNbActive(cubes map[Point]State, p Point) int {
	cnt := 0
	neighs := getNeighbors(p)
	for _, np := range neighs {
		if _, ok := cubes[np]; ok && np != p && cubes[np] == ACTIVE {
			cnt++
		}
	}
	return cnt
}

func main() {
	cubes := make(map[Point]State)
	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, c := range line {
			cubes[newPoint(x, y)] = getState(byte(c))
		}
	}

	// simulate the cycles
	const STEP = 6
	for step := 0; step < STEP; step++ {
		cubes = extendCube(cubes)
		nextCubes := getCopyWithInit(cubes, INACTIVE)
		for p, s := range cubes {
			activeCnt := countNbActive(cubes, p)
			if s == ACTIVE && 2 <= activeCnt && activeCnt <= 3 {
				nextCubes[p] = ACTIVE
			}
			if s == INACTIVE && activeCnt == 3 {
				nextCubes[p] = ACTIVE
			}
		}
		cubes = getCopy(nextCubes)
	}

	result := 0
	for _, v := range cubes {
		if v == ACTIVE {
			result++
		}
	}

	fmt.Println(result)
}
