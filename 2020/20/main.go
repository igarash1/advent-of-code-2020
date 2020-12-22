package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"aoc"
)

const (
	LEN_TILE = 10
)

// Grid for tiles and the whole sea
type Row []byte
type Grid []Row

func (g Grid) print() {
	for y := 0; y < len(g); y++ {
		log.Println(g[y])
	}
}

func (g Grid) copy() Grid {
	dest := make(Grid, len(g))
	for y := 0; y < len(g); y++ {
		dest[y] = make(Row, len(g[y]))
		for x := 0; x < len(g[y]); x++ {
			dest[y][x] = g[y][x]
		}
	}
	return dest
}

func rotateRight(tile Grid) Grid {
	dest := tile.copy()
	n := len(tile)
	for d := 0; d < n/2; d++ {
		for i := d; i < n-d-1; i++ {
			dest[d][i], dest[n-1-i][d],
				dest[n-1-d][n-1-i], dest[i][n-1-d] =
				dest[n-1-i][d], dest[n-1-d][n-1-i], dest[i][n-1-d], tile[d][i]
		}
	}
	return dest
}

func flip(tile Grid) Grid {
	dest := tile.copy()
	n := len(tile)
	for i := 0; i < n; i++ {
		for j := 0; j < (n+1)/2; j++ {
			dest[i][j], dest[i][n-1-j] = dest[i][n-1-j], dest[i][j]
		}
	}
	return dest
}

func checkHorizontal(left Grid, right Grid) bool {
	for y := 0; y < len(left); y++ {
		if left[y][len(left[y])-1] != right[y][0] {
			return false
		}
	}
	return true
}

func checkVertical(up Grid, down Grid) bool {
	for x := 0; x < len(up[0]); x++ {
		if up[len(up)-1][x] != down[0][x] {
			return false
		}
	}
	return true
}

// struct to store state for backtraking
type Solver struct {
	Tiles map[int]Grid
	// the number of tiles horizontally (and vertically)
	LEN_TILE int
	// use them in backtracking
	found     bool
	used      map[int]bool
	cornerMul int
	result    [][]Grid
}

func NewSolver(tiles map[int]Grid) *Solver {
	s := &Solver{
		Tiles:     tiles,
		LEN_TILE:  int(math.Sqrt(float64(len(tiles)))),
		found:     false,
		used:      make(map[int]bool),
		cornerMul: 1,
	}
	s.result = make([][]Grid, s.LEN_TILE)

	for y := 0; y < s.LEN_TILE; y++ {
		s.result[y] = make([]Grid, s.LEN_TILE)
	}
	for y := 0; y < s.LEN_TILE; y++ {
		for x := 0; x < s.LEN_TILE; x++ {
			s.result[y][x] = make(Grid, s.LEN_TILE)
		}
	}

	return s
}

func (s *Solver) valid(y, x int) bool {
	if y > 0 && !checkVertical(s.result[y-1][x], s.result[y][x]) {
		return false
	}
	if x > 0 && !checkHorizontal(s.result[y][x-1], s.result[y][x]) {
		return false
	}
	return true
}

func (s *Solver) backtrack(y, x int) {
	if s.found {
		// abort backtracking
		return
	}

	if y == s.LEN_TILE-1 && x == s.LEN_TILE-1 {
		// s.printIDs()
		fmt.Println("The ansewr to #1: ", s.cornerMul)
		s.findSeaMonsters()
		s.found = true
		return
	}

	for nid := range s.Tiles {
		if s.used[nid] {
			continue
		}
		ny, nx := y+(x+1)/s.LEN_TILE, (x+1)%s.LEN_TILE

		isCorner := (nx == 0 && ny == 0) || (ny == 0 && nx == s.LEN_TILE-1) ||
			(ny == s.LEN_TILE-1 && nx == 0) || (ny == s.LEN_TILE-1 && nx == s.LEN_TILE-1)
		if isCorner {
			s.cornerMul *= nid
		}

		s.result[ny][nx] = s.Tiles[nid].copy()

		s.used[nid] = true
		for d := 0; d < 4; d++ {
			if s.valid(ny, nx) {
				s.backtrack(ny, nx)
			}
			s.result[ny][nx] = flip(s.result[ny][nx])
			if s.valid(ny, nx) {
				s.backtrack(ny, nx)
			}
			s.result[ny][nx] = flip(s.result[ny][nx])
			s.result[ny][nx] = rotateRight(s.result[ny][nx])
		}
		s.used[nid] = false

		if isCorner {
			s.cornerMul /= nid
		}
	}
}

func (s *Solver) findSeaMonsters() {
	// create a sea from the result
	rN := s.LEN_TILE * (LEN_TILE - 2)
	sea := make(Grid, rN)
	for y := 0; y < s.LEN_TILE; y++ {
		for yy := 1; yy < LEN_TILE-1; yy++ {
			for x := 0; x < s.LEN_TILE; x++ {
				sea[y*(LEN_TILE-2)+yy-1] = append(sea[y*(LEN_TILE-2)+yy-1], s.result[y][x][yy][1:LEN_TILE-1]...)
			}
		}
	}

	seaMonster := Grid{Row("                  # "), Row("#    ##    ##    ###"), Row(" #  #  #  #  #  #   ")}

	match := func(sy, sx int) bool {
		for y := range seaMonster {
			for x := range seaMonster[y] {
				ny, nx := y+sy, x+sx
				if seaMonster[y][x] == '#' && sea[ny][nx] == '.' {
					return false
				}
			}
		}
		return true
	}

	fill := func(sy, sx int) {
		for y := range seaMonster {
			for x := range seaMonster[y] {
				ny, nx := y+sy, x+sx
				if seaMonster[y][x] == '#' && sea[ny][nx] != '.' {
					sea[ny][nx] = 'O'
				}
			}
		}
	}

	check := func() {
		for sy := 0; sy < rN-len(seaMonster); sy++ {
			for sx := 0; sx < rN-len(seaMonster[0]); sx++ {
				if match(sy, sx) {
					fill(sy, sx)
				}
			}
		}
	}

	for d := 0; d < 4; d++ {
		check()
		sea = flip(sea)
		check()
		sea = flip(sea)
		sea = rotateRight(sea)
	}

	countRough := 0
	for y := range sea {
		for x := range sea[y] {
			if sea[y][x] == '#' {
				countRough++
			}
		}
	}

	fmt.Println("The ansewr to #2: ", countRough)
}

func main() {
	tiles := make(map[int]Grid)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		tileID := aoc.ToInt(line[5 : len(line)-1])
		for h := 0; h < LEN_TILE; h++ {
			scanner.Scan()
			tiles[tileID] = append(tiles[tileID], []byte(scanner.Text()))
		}

		if !scanner.Scan() {
			break
		}
	}
	log.Println("#tiles = ", len(tiles))

	s := NewSolver(tiles)
	s.backtrack(0, -1)
}
