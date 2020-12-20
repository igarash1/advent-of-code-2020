package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	"../helper"
)

const (
	NTILE = 10
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
type State struct {
	Tiles map[int]Grid
	NSea  int

	// use them in backtracking
	found     bool
	used      map[int]bool
	cornerMul int
	result    [][]Grid
}

func (s *State) Init() {
	s.result = make([][]Grid, s.NSea)
	for y := 0; y < s.NSea; y++ {
	for y := 0; y < s.NSea; y++ {
		s.result[y] = make([]Grid, s.NSea)
	}
	for y := 0; y < s.NSea; y++ {
		for x := 0; x < s.NSea; x++ {
			s.result[y][x] = make(Grid, s.NSea)
		}
	}

	s.used = make(map[int]bool)
	s.cornerMul = 1
	s.found = false
}

func (s *State) valid(y, x int) bool {
	if y > 0 && !checkVertical(s.result[y-1][x], s.result[y][x]) {
		return false
	}
	if x > 0 && !checkHorizontal(s.result[y][x-1], s.result[y][x]) {
		return false
	}
	return true
}

func (s *State) dfs(y, x int) {
	if s.found {
		// abort backtracking
		return
	}

	if y == s.NSea-1 && x == s.NSea-1 {
		// s.printIDs()
		fmt.Println("The ansewr to #1: ", s.cornerMul)
		s.countSeaMonster()
		s.found = true
		return
	}

	for nid := range s.Tiles {
		if s.used[nid] {
			continue
		}
		ny, nx := y+(x+1)/s.NSea, (x+1)%s.NSea

		isCorner := (nx == 0 && ny == 0) || (ny == 0 && nx == s.NSea-1) ||
			(ny == s.NSea-1 && nx == 0) || (ny == s.NSea-1 && nx == s.NSea-1)
		if isCorner {
			s.cornerMul *= nid
		}

		s.result[ny][nx] = s.Tiles[nid].copy()

		s.used[nid] = true
		for d := 0; d < 4; d++ {
			if s.valid(ny, nx) {
				s.dfs(ny, nx)
			}
			s.result[ny][nx] = flip(s.result[ny][nx])
			if s.valid(ny, nx) {
				s.dfs(ny, nx)
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

func (s *State) countSeaMonster() {
	// create a sea from the result
	rN := s.NSea * (NTILE - 2)
	sea := make(Grid, rN)
	for y := 0; y < s.NSea; y++ {
		for yy := 1; yy < NTILE-1; yy++ {
			for x := 0; x < s.NSea; x++ {
				sea[y*(NTILE-2)+yy-1] = append(sea[y*(NTILE-2)+yy-1], s.result[y][x][yy][1:NTILE-1]...)
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
	s := State{Tiles: make(map[int]Grid)}

	scanner := bufio.NewScanner(os.Stdin)
	tileNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		tileID := helper.ToInt(line[5 : len(line)-1])
		for h := 0; h < NTILE; h++ {
			scanner.Scan()
			s.Tiles[tileID] = append(s.Tiles[tileID], []byte(scanner.Text()))
		}

		tileNumber++
		if !scanner.Scan() {
			break
		}
	}
	log.Println("#tiles = ", tileNumber)

	s.NSea = int(math.Sqrt(float64(tileNumber)))

	s.Init()
	s.dfs(0, -1)
}
