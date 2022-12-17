package main

import (
	"bufio"
	"strings"
)

func main() {
}

type coordinate struct {
	x int
	y int
}

func newCoordinate(x, y int) coordinate {
	return coordinate{x, y}
}

type rows [][]rune

func (i rows) find(in rune) *coordinate {
	for y, row := range i {
		for x, cell := range row {
			if cell == in {
				c := newCoordinate(x, y)
				return &c
			}
		}
	}

	return nil
}

func newRows(in string) rows {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		out     = rows{}
	)
	for scanner.Scan() {
		var (
			cs = strings.TrimSpace(scanner.Text())
			r  = []rune{}
		)
		for _, c := range cs {
			r = append(r, c)
		}
		out = append(out, r)
	}

	return out
}

type move func(i grid, a, b coordinate) bool

type end func(i grid, a, b coordinate) bool

type grid struct {
	rows rows
	end  end
	move move
}

func newGrid(rows rows, move move, end end) *grid {
	return &grid{
		rows: rows,
		move: move,
		end:  end,
	}
}

func (i grid) bfs(start coordinate) int {
	var (
		visited = make(map[coordinate]int)
		queue   = []coordinate{}
		curr    coordinate
	)

	queue = append(queue, start)
	visited[start] = 0

	for len(queue) > 0 {
		curr = queue[0]
		queue = queue[1:]

		var (
			value = visited[curr]
			next  = i.getAdjacent(curr)
		)

		for _, n := range next {
			if i.end != nil && i.end(i, curr, n) {
				return value + 1
			}

			if _, ok := visited[n]; ok {
				continue
			}

			visited[n] = value + 1
			queue = append(queue, n)

		}
	}

	return -1
}

func (i grid) getAdjacent(c coordinate) []coordinate {
	var (
		left  = newCoordinate(c.x-1, c.y)
		right = newCoordinate(c.x+1, c.y)
		up    = newCoordinate(c.x, c.y-1)
		down  = newCoordinate(c.x, c.y+1)
		out   = []coordinate{}
	)

	if i.move == nil {
		return out
	}

	if ok := i.move(i, c, left); ok {
		out = append(out, left)
	}
	if ok := i.move(i, c, right); ok {
		out = append(out, right)
	}
	if ok := i.move(i, c, up); ok {
		out = append(out, up)
	}
	if ok := i.move(i, c, down); ok {
		out = append(out, down)
	}

	return out
}

func (i grid) getValue(c coordinate) (rune, bool) {
	if c.y < 0 || c.y >= len(i.rows) || c.x < 0 || c.x >= len(i.rows[c.y]) {
		return -1, false
	}

	return i.rows[c.y][c.x], true
}

func exercise1(stream string) int {
	var (
		rows = newRows(stream)
		end  = func(i grid, a, b coordinate) bool {
			av, _ := i.getValue(a)
			bv, _ := i.getValue(b)

			return av == 'z' && bv == 'E'
		}
		move = func(i grid, a, b coordinate) bool {
			av, _ := i.getValue(a)
			bv, ok := i.getValue(b)
			if !ok {
				return false
			}

			if av == 'S' && bv == 'a' || av == 'z' && bv == 'E' {
				return true
			}

			return bv-av <= 1
		}
		grid = newGrid(rows, move, end)
	)

	start := rows.find('S')
	if start == nil {
		panic("could not find start")
	}

	return grid.bfs(*start)
}

func exercise2(stream string) int {
	var (
		rows = newRows(stream)
		end  = func(i grid, a, b coordinate) bool {
			av, _ := i.getValue(a)
			bv, _ := i.getValue(b)

			return av == 'b' && bv == 'a' || av == 'b' && bv == 'S'
		}
		move = func(i grid, a, b coordinate) bool {
			av, _ := i.getValue(a)
			bv, ok := i.getValue(b)
			if !ok {
				return false
			}

			if av == 'E' {
				return bv == 'z'
			}

			if bv == 'S' {
				return av == 'a' || av == 'b'
			}

			return av-bv <= 1
		}
		grid = newGrid(rows, move, end)
	)

	start := rows.find('E')
	if start == nil {
		panic("could not find start")
	}

	return grid.bfs(*start)
}
