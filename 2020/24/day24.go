package day24

import (
	"aoc"
	//"log"
	"strings"
)

type Point struct {
	x, y int
}

type HexTiles struct {
	dx    [2][6]int
	dy    [2][6]int
	DIR   map[string]int
	blackTiles map[Point]bool
}

func newHexTiles() *HexTiles {
	hex := &HexTiles{
		dx:  [2][6]int{{1, 1, 0, -1, -1, 0}, {1, 1, 0, -1, -1, 0}},
		dy:  [2][6]int{{0, -1, -1, -1, 0, 1}, {1, 0, -1, 0, 1, 1}},
		DIR: map[string]int{"e": 0, "se": 1, "sw": 2, "w": 3, "nw": 4, "ne": 5}}
	hex.blackTiles = make(map[Point]bool)
	return hex
}

func (hex *HexTiles) refPoint() Point {
	return Point{0, 0}
}

func (hex *HexTiles) flip(p Point) {
	if v, ok := hex.blackTiles[p]; ok && v {
		delete(hex.blackTiles, p)
	} else {
		hex.blackTiles[p] = true
	}
}

func (hex *HexTiles) move(p Point, dir string) Point {
	u := aoc.Abs(p.x) % 2
	r := hex.DIR[dir]
	p.x += hex.dx[u][r]
	p.y += hex.dy[u][r]
	return p
}

func (hex *HexTiles) dirExists(dir string) bool {
	_, exist := hex.DIR[dir]
	return exist
}

func (hex *HexTiles) countTotalBlacks() int {
	//log.Println("#black = ", len(tile.blackTiles))
	return len(hex.blackTiles)
}

func part1(input string) (int, *HexTiles) {
	lines := strings.Split(input, "\n")
	hex := newHexTiles()
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		cur := hex.refPoint()
		for i := 0; i < len(line); {
			if i+2 <= len(line) && hex.dirExists(line[i:i+2]) {
				cur = hex.move(cur, line[i:i+2])
				i += 2
			} else {
				cur = hex.move(cur, line[i:i+1])
				i++
			}
		}
		hex.flip(cur)
	}
	return hex.countTotalBlacks(), hex
}

func (hex *HexTiles) isBlack(p Point) bool {
	_, exist := hex.blackTiles[p]
	return exist
}

func (hex *HexTiles) getBlacks() []Point {
	var blk []Point
	for p := range hex.blackTiles {
		blk = append(blk, p)
	}
	return blk
}

func (hex *HexTiles) getNeighs(p Point) []Point {
	var neighs []Point
	for r := range hex.DIR {
		neighs = append(neighs, hex.move(p, r))
	}
	return neighs
}

func (hex *HexTiles) getCandidateWhites() []Point {
	blacks := hex.getBlacks()
	candWhiteMap := make(map[Point]bool)
	for _, p := range blacks {
		neighs := hex.getNeighs(p)
		for _, p := range neighs {
			if !hex.isBlack(p) {
				candWhiteMap[p] = true
			}
		}
	}
	var candWhiteList []Point
	for p, _ := range candWhiteMap {
		candWhiteList = append(candWhiteList, p)
	}
	return candWhiteList
}

func (hex *HexTiles) countNeighB(p Point) int {
	neighs := hex.getNeighs(p)
	cnt := 0
	for _, p := range neighs {
		if hex.isBlack(p) {
			cnt++
		}
	}
	return cnt
}

func part2(input string) int {
	_, tile := part1(input)
	for d := 0; d < 100; d++ {
		// defer flipping to apply the rules simultaneously
		deferredFlips := make(map[Point]bool)
		// determine which blackTiles to be flipped
		blacks := tile.getBlacks()
		for _, p := range blacks {
			cnt := tile.countNeighB(p)
			if cnt == 0 || 2 < cnt {
				deferredFlips[p] = true
			}
		}
		whites := tile.getCandidateWhites()
		for _, p := range whites {
			cnt := tile.countNeighB(p)
			if cnt == 2 {
				deferredFlips[p] = true
			}
		}
		// update the blackTiles
		for p := range deferredFlips {
			tile.flip(p)
		}
		//log.Printf("Day %d: %d\n", d + 1, tile.countTotalBlacks())
	}
	return tile.countTotalBlacks()
}
