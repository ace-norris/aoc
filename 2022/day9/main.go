package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {}

type direction int64

func newDirection(in string) direction {
	switch in {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	default:
		panic(fmt.Sprintf("unsupported direction: %s", in))
	}
}

const (
	Up direction = iota
	Down
	Left
	Right
)

type instruction struct {
	Direction direction
	Distance  int
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

		d, err := strconv.Atoi(exp[2])
		if err != nil {
			continue
		}

		out = append(out, instruction{
			Direction: newDirection(exp[1]),
			Distance:  d,
		})
	}

	return out
}

type position struct {
	X, Y int
}

func (i position) string() string {
	return fmt.Sprintf("%v:%v", i.X, i.Y)
}

func (i position) adjacent(in position) bool {
	if i.X == in.X && i.Y == in.Y {
		// . . .
		// . S .
		// . . .
		return true
	}

	if i.Y == in.Y && (i.X+1 == in.X || i.X-1 == in.X) {
		// . . .
		// . H T
		// . . .

		// . . .
		// T H .
		// . . .
		return true
	}

	if i.Y-1 == in.Y && (i.X+1 == in.X || i.X-1 == in.X) {
		// . . .
		// . H .
		// T . .

		// . . .
		// . H .
		// . . T
		return true
	}

	if i.Y+1 == in.Y && (i.X+1 == in.X || i.X-1 == in.X) {
		// T . .
		// . H .
		// . . .

		// . . T
		// . H .
		// . . .
		return true
	}

	if i.X == in.X && (i.Y+1 == in.Y || i.Y-1 == in.Y) {
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

func (i *position) move(in direction) {
	switch in {
	case Up:
		i.Y += 1
	case Down:
		i.Y -= 1
	case Left:
		i.X -= 1
	case Right:
		i.X += 1
	}
}

func (i *position) follow(in position) {
	if i.adjacent(in) {
		return
	}

	if i.X-2 == in.X && i.Y == in.Y {
		i.move(Left)
	}
	if i.X+2 == in.X && i.Y == in.Y {
		i.move(Right)
	}
	if i.Y+2 == in.Y && i.X == in.X {
		i.move(Up)
	}
	if i.Y-2 == in.Y && i.X == in.X {
		i.move(Down)
	}

	if i.adjacent(in) {
		return
	}

	if i.Y > in.Y {
		i.move(Down)
	}

	if i.Y < in.Y {
		i.move(Up)
	}

	if i.X > in.X {
		i.move(Left)
	}

	if i.X < in.X {
		i.move(Right)
	}
}

func (i position) clone() position {
	return newPosition(i.X, i.Y)
}

func newPosition(x, y int) position {
	return position{x, y}
}

type positions []position

func (i positions) unique() []position {
	var (
		up  = map[string]position{}
		out = []position{}
	)

	for _, v := range i {
		up[v.string()] = v
	}

	for _, v := range up {
		out = append(out, v)
	}

	return out
}

func (i positions) string() string {
	var (
		minx = 0
		maxx = 0
		miny = 0
		maxy = 0
		up   = map[string]position{}
	)

	for _, p := range i {
		up[p.string()] = p

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
			if _, ok := up[newPosition(x, y).string()]; ok {
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

type grid struct {
	TailLength    int
	HeadPositions positions
	TailPositions []positions
}

func (i *grid) instruct(in instruction) {
	for d := 0; d < in.Distance; d++ {
		chp := i.HeadPositions[len(i.HeadPositions)-1]
		ctps := i.TailPositions[len(i.TailPositions)-1]

		// new head position
		nhp := chp.clone()
		nhp.move(in.Direction)
		i.HeadPositions = append(i.HeadPositions, nhp)

		// new tail positions
		ntps := []position{}
		for y, ctp := range ctps {
			ptp := nhp.clone()
			if len(ntps) > 0 {
				ptp = ntps[y-1]
			}

			ntp := ctp.clone()
			ntp.follow(ptp)
			ntps = append(ntps, ntp)
		}
		i.TailPositions = append(i.TailPositions, ntps)
	}
}

func (i grid) extractTailPositions(in int) positions {
	out := positions{}

	for _, v := range i.TailPositions {
		out = append(out, v[in])
	}

	return out
}

func newGrid(tailLength int) grid {
	tp := positions{}
	for i := 0; i < tailLength; i++ {
		tp = append(tp, newPosition(0, 0))
	}

	return grid{
		TailLength:    tailLength,
		HeadPositions: []position{newPosition(0, 0)},
		TailPositions: []positions{tp},
	}
}

func process(stream string, tailLength int) (int, string) {
	var (
		instructions = newInstructions(stream)
		grid         = newGrid(tailLength)
	)
	for _, v := range instructions {
		grid.instruct(v)
	}

	tp := grid.extractTailPositions(grid.TailLength - 1)
	return len(tp.unique()), tp.string()
}
