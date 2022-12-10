package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {}

type instruction struct {
	Direction string
	Count     int
}

type instructions []instruction

func newInstructions(in string) instructions {
	scanner := bufio.NewScanner(strings.NewReader(in))

	var (
		out = instructions{}
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		exp := regexp.MustCompile(`^([UDLR])\s(\d+)$`).FindStringSubmatch(line)
		if len(exp) != 3 {
			continue
		}

		c, err := strconv.Atoi(exp[2])
		if err != nil {
			continue
		}

		out = append(out, instruction{
			Direction: exp[1],
			Count:     c,
		})
	}

	return out
}

type position struct {
	X, Y int
}

func (i position) key() string {
	return fmt.Sprintf("%v:%v", i.X, i.Y)
}

func (i position) adjacent(p position) bool {
	if i.X == p.X && i.Y == p.Y {
		// . . .
		// . S .
		// . . .
		return true
	}

	if i.Y == p.Y && (i.X+1 == p.X || i.X-1 == p.X) {
		// . . .
		// . H T
		// . . .

		// . . .
		// T H .
		// . . .
		return true
	}

	if i.Y-1 == p.Y && (i.X+1 == p.X || i.X-1 == p.X) {
		// . . .
		// . H .
		// T . .

		// . . .
		// . H .
		// . . T
		return true
	}

	if i.Y+1 == p.Y && (i.X+1 == p.X || i.X-1 == p.X) {
		// T . .
		// . H .
		// . . .

		// . . T
		// . H .
		// . . .
		return true
	}

	if i.X == p.X && (i.Y+1 == p.Y || i.Y-1 == p.Y) {
		// . T .
		// . H .
		// . . .

		// . . .
		// . H .
		// . T .
		return true
	}

	return false
}

func (i position) move(direction string) position {
	switch direction {
	case "U":
		return newPosition(i.X, i.Y+1)
	case "D":
		return newPosition(i.X, i.Y-1)
	case "L":
		return newPosition(i.X-1, i.Y)
	case "R":
		return newPosition(i.X+1, i.Y)
	default:
		return i.clone()
	}
}

func (i position) follow(in position) position {
	n := i.clone()
	if n.adjacent(in) {
		return n
	}

	if i.X-2 == in.X && i.Y == in.Y {
		n = i.move("L")
	}
	if i.X+2 == in.X && i.Y == in.Y {
		n = i.move("R")
	}
	if i.Y+2 == in.Y && i.X == in.X {
		n = i.move("U")
	}
	if i.Y-2 == in.Y && i.X == in.X {
		n = i.move("D")
	}

	if n.adjacent(in) {
		return n
	}

	if n.Y > in.Y {
		n = n.move("D")
	}

	if n.Y < in.Y {
		n = n.move("U")
	}

	if n.X > in.X {
		n = n.move("L")
	}

	if n.X < in.X {
		n = n.move("R")
	}

	return n
}

func (i position) clone() position {
	return newPosition(i.X, i.Y)
}

func newPosition(x, y int) position {
	return position{x, y}
}

type grid struct {
	TailLength    int
	HeadPositions []position
	TailPositions [][]position
}

func (i *grid) applyInstruction(instruction instruction) {
	for j := 0; j < instruction.Count; j++ {
		chp := i.HeadPositions[len(i.HeadPositions)-1]
		ctps := i.TailPositions[len(i.TailPositions)-1]
		nhp := chp.move(instruction.Direction)

		i.HeadPositions = append(i.HeadPositions, nhp)
		i.TailPositions = append(i.TailPositions, i.nextTailPositions(nhp, ctps))
	}
}

func (i grid) nextTailPositions(head position, ctps []position) []position {
	out := []position{}
	for j, p := range ctps {
		pp := head.clone()
		if len(out) > 0 {
			pp = out[j-1]
		}

		out = append(out, p.follow(pp))
	}

	return out
}

func (i grid) extractTailPositions(in int) []position {
	out := []position{}

	for _, v := range i.TailPositions {
		out = append(out, v[in])
	}

	return out
}

func (i grid) uniquePositions(positions []position) []position {
	var (
		up  = map[string]position{}
		out = []position{}
	)

	for _, v := range positions {
		up[v.key()] = v
	}

	for _, v := range up {
		out = append(out, v)
	}

	return out
}

func (i grid) string(positions []position) string {
	var (
		minx = 0
		maxx = 0
		miny = 0
		maxy = 0
		up   = map[string]position{}
	)

	for _, p := range positions {
		up[p.key()] = p

		if p.X > maxx {
			maxx = p.X
		}
		if p.X < minx {
			minx = p.X
		}
		if p.Y > maxy {
			maxy = p.Y
		}
		if p.Y < miny {
			miny = p.Y
		}
	}

	out := ""
	for y := maxy; y >= miny; y-- {
		for x := minx; x <= maxx; x++ {
			if _, ok := up[newPosition(x, y).key()]; ok {
				if x == 0 && y == 0 {
					out += "S"
				} else {
					out += "#"
				}
			} else {
				out += "."
			}
		}
		out += "\n"
	}

	return out
}

func newGrid(tailLength int) grid {
	tp := []position{}
	for i := 0; i < tailLength; i++ {
		tp = append(tp, newPosition(0, 0))
	}

	return grid{
		TailLength:    tailLength,
		HeadPositions: []position{newPosition(0, 0)},
		TailPositions: [][]position{tp},
	}
}

func exercise1(in string) (int, string) {
	var (
		instructions = newInstructions(in)
		grid         = newGrid(1)
	)
	for _, instruction := range instructions {
		grid.applyInstruction(instruction)
	}

	tp := grid.extractTailPositions(grid.TailLength - 1)
	return len(grid.uniquePositions(tp)), grid.string(tp)
}

func exercise2(in string) (int, string) {
	var (
		instructions = newInstructions(in)
		grid         = newGrid(9)
	)
	for _, instruction := range instructions {
		grid.applyInstruction(instruction)
	}

	tp := grid.extractTailPositions(grid.TailLength - 1)
	return len(grid.uniquePositions(tp)), grid.string(tp)
}
