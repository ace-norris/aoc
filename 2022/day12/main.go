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

func (i coordinate) getValue(rows [][]rune) (rune, bool) {
	if i.x < 0 || i.x >= len(rows) || i.y < 0 || i.y >= len(rows[i.x]) {
		return -1, false
	}

	return rows[i.x][i.y], true
}

func (i coordinate) canMove(rows [][]rune, to coordinate) bool {
	fv, _ := i.getValue(rows)

	tv, ok := to.getValue(rows)
	if !ok {
		return false
	}

	return tv-fv <= 1
}

func (i coordinate) getAdjacent(rows [][]rune) []coordinate {
	var (
		left  = coordinate{i.x - 1, i.y}
		right = coordinate{i.x + 1, i.y}
		up    = coordinate{i.x, i.y - 1}
		down  = coordinate{i.x, i.y + 1}
		out   = []coordinate{}
	)

	if ok := i.canMove(rows, left); ok {
		out = append(out, left)
	}
	if ok := i.canMove(rows, right); ok {
		out = append(out, right)
	}
	if ok := i.canMove(rows, up); ok {
		out = append(out, up)
	}
	if ok := i.canMove(rows, down); ok {
		out = append(out, down)
	}

	return out
}

func (i coordinate) isEnd(end coordinate) bool {
	return i.x == end.x && i.y == end.y
}

type grid struct {
	rows  [][]rune
	start coordinate
	end   coordinate
}

func newGrid(in string, start, first, last, end rune) *grid {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		y       = 0
		x       = 0
		out     = &grid{
			rows: [][]rune{},
		}
	)
	for scanner.Scan() {
		y = 0

		var (
			cs = strings.TrimSpace(scanner.Text())
			r  = []rune{}
		)
		for _, c := range cs {
			if c == start {
				out.start = coordinate{x, y}
				c = first
			}
			if c == end {
				out.end = coordinate{x, y}
				c = last
			}

			r = append(r, c)
			y++
		}
		out.rows = append(out.rows, r)
		x++
	}

	return out
}

func (i grid) bfs() int {
	var (
		visited = make(map[coordinate]int)
		queue   = []coordinate{}
		curr    coordinate
	)

	queue = append(queue, i.start)
	visited[i.start] = 0

	for len(queue) > 0 {
		curr = queue[0]
		queue = queue[1:]

		var (
			value = visited[curr]
			next  = curr.getAdjacent(i.rows)
		)

		for _, n := range next {
			if n.isEnd(i.end) {
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

func exercise1(stream string) int {
	grid := newGrid(stream, 'S', 'a', 'z', 'E')

	return grid.bfs()
}

func exercise2(in string) int {
	scanner := bufio.NewScanner(strings.NewReader(in))

	var (
		total = 0
	)
	for scanner.Scan() {
		// line := strings.TrimSpace(scanner.Text())
		total++
	}

	return total
}
