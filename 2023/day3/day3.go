package day3

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func Exercise1(stream string) int {
	grid := newGrid(stream)

	num := grid.getNumbers()
	sym := grid.getSymbols()

	out := make(numbers, 0)
	for _, s := range sym {
		for _, n := range num {
			if n.isAdjacent(s) {
				out = append(out, n)
			}
		}
	}

	return out.sum()
}

func Exercise2(stream string) int {
	grid := newGrid(stream)

	num := grid.getNumbers()
	gears := grid.getGears()

	out := 0
	for _, g := range gears {

		adj := make(numbers, 0)

		for _, n := range num {
			if n.isAdjacent(g) {
				adj = append(adj, n)
			}
		}

		if len(adj) == 2 {
			out += adj[0].value() * adj[1].value()
		}
	}

	return out
}

type grid []cells

func newGrid(in string) grid {
	var (
		scanner = bufio.NewScanner(strings.NewReader(in))
		out     = make(grid, 0)
		y       = 0
	)
	for scanner.Scan() {
		row := make(cells, 0)
		for x, v := range strings.Split(strings.TrimSpace(scanner.Text()), "") {
			row = append(row, cell{
				value: v,
				x:     x,
				y:     y,
			})
		}

		out = append(out, row)
		y++
	}

	return out
}

func (i grid) getNumbers() numbers {
	out := make(numbers, 0)

	for _, r := range i {
		b := make(number, 0)

		for _, c := range r {
			if c.isNumber() {
				b = append(b, c)
			} else {
				out = append(out, b)
				b = make(number, 0)
			}
		}
		if len(b) > 0 {
			out = append(out, b)
		}
	}

	return out
}

func (i grid) getGears() cells {
	out := make(cells, 0)

	for _, r := range i {
		for _, c := range r {
			if c.isGear() {
				out = append(out, c)
			}
		}
	}

	return out
}

func (i grid) getSymbols() cells {
	out := make(cells, 0)

	for _, r := range i {
		for _, c := range r {
			if c.isSymbol() {
				out = append(out, c)
			}
		}
	}

	return out
}

type cells []cell

type cell struct {
	value string
	x, y  int
}

func (i cell) isSymbol() bool {
	return !i.isNumber() && !i.isPeriod()
}

func (i cell) isNumber() bool {
	return len(regexp.MustCompile("(\\d)").FindStringSubmatch(i.value)) == 2
}

func (i cell) isPeriod() bool {
	return i.value == `.`
}

func (i cell) isGear() bool {
	return i.value == `*`
}

func (c cell) isAdjacent(cell cell) bool {
	var (
		f = c.x + 1
		b = c.x - 1
		u = c.y - 1
		d = c.y + 1
	)
	// up
	if cell.x == c.x && cell.y == u {
		return true
	}

	// up/back
	if cell.x == b && cell.y == u {
		return true
	}

	// up/forward
	if cell.x == f && cell.y == u {
		return true
	}

	// back
	if cell.x == b && cell.y == c.y {
		return true
	}

	// forward
	if cell.x == f && cell.y == c.y {
		return true
	}

	// down/back
	if cell.x == b && cell.y == d {
		return true
	}

	// down/forward
	if cell.x == f && cell.y == d {
		return true
	}

	if cell.x == c.x && cell.y == d {
		return true
	}

	return false
}

type numbers []number

func (i numbers) sum() int {
	out := 0

	for _, n := range i {
		out += n.value()
	}

	return out
}

type number []cell

func (i number) value() int {
	s := ""
	for _, y := range i {
		s += y.value
	}
	if s == "" {
		return 0
	}

	n, _ := strconv.Atoi(s)

	return n
}

func (i number) isAdjacent(cell cell) bool {
	for _, c := range i {
		if c.isAdjacent(cell) {
			return true
		}
	}

	return false
}
