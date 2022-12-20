package grid

import (
	"bufio"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{x, y}
}

type Rows [][]rune

func (i Rows) Find(in rune) *Coordinate {
	for y, row := range i {
		for x, cell := range row {
			if cell == in {
				c := NewCoordinate(x, y)
				return &c
			}
		}
	}

	return nil
}

func NewRows(in string) Rows {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		out     = Rows{}
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

type Move func(a Coordinate, av rune, b Coordinate, bv rune) bool

type End func(a Coordinate, av rune, b Coordinate, bv rune) bool

type Grid struct {
	rows  Rows
	start Coordinate
	end   End
	move  Move
}

func NewGrid(rows Rows) Grid {
	return Grid{
		rows: rows,
	}
}

func (i *Grid) SetStart(in Coordinate) {
	i.start = in
}

func (i *Grid) SetMove(in Move) {
	i.move = in
}

func (i *Grid) SetEnd(in End) {
	i.end = in
}

func (i Grid) BFS() int {
	var (
		visited = make(map[Coordinate]int)
		queue   = []Coordinate{}
	)

	queue = append(queue, i.start)
	visited[i.start] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		var (
			value = visited[curr]
			next  = i.getAdjacent(curr)
		)

		for _, n := range next {
			if i.tryEnd(curr, n) {
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

func (i Grid) getAdjacent(c Coordinate) []Coordinate {
	var (
		left  = NewCoordinate(c.x-1, c.y)
		right = NewCoordinate(c.x+1, c.y)
		up    = NewCoordinate(c.x, c.y-1)
		down  = NewCoordinate(c.x, c.y+1)
		out   = []Coordinate{}
	)

	if i.move == nil {
		return out
	}

	if ok := i.tryMove(c, left); ok {
		out = append(out, left)
	}
	if ok := i.tryMove(c, right); ok {
		out = append(out, right)
	}
	if ok := i.tryMove(c, up); ok {
		out = append(out, up)
	}
	if ok := i.tryMove(c, down); ok {
		out = append(out, down)
	}

	return out
}

func (i Grid) getValue(c Coordinate) (rune, bool) {
	if c.y < 0 || c.y >= len(i.rows) || c.x < 0 || c.x >= len(i.rows[c.y]) {
		return -1, false
	}

	return i.rows[c.y][c.x], true
}

func (i Grid) tryMove(a, b Coordinate) bool {
	if i.move == nil {
		return false
	}

	av, _ := i.getValue(a)
	bv, ok := i.getValue(b)
	if !ok {
		return false
	}

	return i.move(a, av, b, bv)
}

func (i Grid) tryEnd(a, b Coordinate) bool {
	if i.end == nil {
		return false
	}

	av, _ := i.getValue(a)
	bv, ok := i.getValue(b)
	if !ok {
		return false
	}

	return i.end(a, av, b, bv)
}
