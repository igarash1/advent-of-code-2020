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

type Point struct {
	x, y, z int
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

func extendCube(cubes map[Point]State) map[Point]State {
	for p := range cubes {
		for rx := -1; rx <= 1; rx++ {
			for ry := -1; ry <= 1; ry++ {
				for rz := -1; rz <= 1; rz++ {
					np := Point{p.x + rx, p.y + ry, p.z + rz}
					if _, ok := cubes[np]; !ok {
						cubes[np] = INACTIVE
					}
				}
			}
		}
	}
	return cubes
}

func countNbActive(cubes map[Point]State, p Point) int {
	cnt := 0
	for rx := -1; rx <= 1; rx++ {
		for ry := -1; ry <= 1; ry++ {
			for rz := -1; rz <= 1; rz++ {
				np := Point{p.x + rx, p.y + ry, p.z + rz}
				if _, ok := cubes[np]; ok && np != p && cubes[np] == ACTIVE {
					cnt++
				}
			}
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
			cubes[Point{x, y, 0}] = getState(byte(c))
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
